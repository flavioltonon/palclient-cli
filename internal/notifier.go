package internal

import (
	"fmt"
	"palclient-cli/pkg/discord"
	"palclient-cli/pkg/palserver"
	"time"
)

type Notifier struct {
	discordClient   *discord.Client
	palserverClient *palserver.Client
}

func NewNotifier(discordClient *discord.Client, palserverClient *palserver.Client) *Notifier {
	return &Notifier{
		discordClient:   discordClient,
		palserverClient: palserverClient,
	}
}

func (s *Notifier) NotifyServerRestart(wait time.Duration) error {
	message := fmt.Sprintf("The server will restart in %0.f minutes.", wait.Minutes())

	if wait == 0 {
		message = "The server will restart immediately."
	}

	if err := s.palserverClient.Broadcast(message); err != nil {
		return fmt.Errorf("broadcasting to palserver: %w", err)
	}

	if err := s.discordClient.ExecuteWebhook(message); err != nil {
		return fmt.Errorf("executing Discord webhook: %w", err)
	}

	return nil
}
