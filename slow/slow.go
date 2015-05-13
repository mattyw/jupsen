package slow

import (
	"errors"
	"log"
	"os/exec"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
slow puts a delay on eth0 for the specified unit

Examples:

juju jupsen slow wordpress/0
`

type Command struct {
	cmd.CommandBase
	unit string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "slow",
		Args:    "unit",
		Purpose: "slow puts a delay on eth0 for the specified unit",
		Doc:     commandDoc,
	}
}

func (c *Command) Init(args []string) error {
	if len(args) < 1 {
		return errors.New("expected a unit name")
	}
	c.unit = args[0]
	return cmd.CheckEmpty(args[1:])
}

func (c *Command) Run(ctx *cmd.Context) error {
	args := []string{"ssh", c.unit}
	args = append(args, common.Slow()...)
	log.Println(args)
	_, err := exec.Command("juju", args...).Output()
	if err != nil {
		return err
	}
	return nil
}
