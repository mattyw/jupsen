package fix

import (
	"errors"
	"github.com/juju/cmd"
	"log"
	"os/exec"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
fix removes any slow or flaky settings jupsen has placed on a unit

Examples:

juju jupsen fix wordpress/0
`

type Command struct {
	cmd.CommandBase
	unit string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "fix",
		Args:    "unit",
		Purpose: "fix removes any slow or flaky setting jupsen has placed on the unit",
		Doc:     commandDoc,
	}
}

func (c *Command) Init(args []string) error {
	if len(args) < 1 {
		return errors.New("expected a unit names")
	}
	c.unit = args[0]
	return cmd.CheckEmpty(args[1:])
}

func (c *Command) Run(ctx *cmd.Context) error {
	args := []string{"ssh", c.unit}
	args = append(args, common.Fix()...)
	log.Println(args)
	_, err := exec.Command("juju", args...).Output()
	if err != nil {
		return err
	}
	return nil
}
