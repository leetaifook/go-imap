package commands

import (
	imap "github.com/emersion/imap/common"
)

// A DELETE command.
// See https://tools.ietf.org/html/rfc3501#section-6.3.3
type Delete struct {
	Mailbox string
}

func (c *Delete) Command() *imap.Command {
	return &imap.Command{
		Name: imap.Delete,
		Arguments: []interface{}{c.Mailbox},
	}
}
