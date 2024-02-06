package palserver

import (
	"fmt"
	"net"
	"strings"

	"github.com/gorcon/rcon"
)

type Client struct {
	connection *rcon.Conn
}

func NewClient(host string, port string, password string) (*Client, error) {
	address := net.JoinHostPort(host, port)

	connection, err := rcon.Dial(address, password)
	if err != nil {
		return nil, fmt.Errorf("dialing server: %w", err)
	}

	return &Client{connection: connection}, nil
}

func (c *Client) Broadcast(message string) error {
	// Replace all whitespaces with underscores so the message doesn't get cut off.
	message = strings.ReplaceAll(message, " ", "_")

	if _, err := c.connection.Execute(fmt.Sprintf("Broadcast %s", message)); err != nil {
		return fmt.Errorf("executing command: %w", err)
	}

	return nil
}
