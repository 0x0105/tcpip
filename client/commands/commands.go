package commands

import (
	"github.com/spf13/cobra"
)

type cmder interface {
	getCommand() *cobra.Command
}

type baseCmd struct {
	cmd *cobra.Command
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}
