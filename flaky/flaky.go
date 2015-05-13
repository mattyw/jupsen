package flaky

import (
	"errors"
	"log"
	"os/exec"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
flaky causes the unit to drop packets

Examples:

juju jupsen flaky wordpress/0
`

type Command struct {
	cmd.CommandBase
	unit string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "flaky",
		Args:    "unit",
		Purpose: "flaky causes the unit to drop packets",
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
	args = append(args, common.Flaky()...)
	log.Println(args)
	_, err := exec.Command("juju", args...).Output()
	if err != nil {
		log.Fatalf("failed to make flaky: %v", err)
	}
	return nil
}
