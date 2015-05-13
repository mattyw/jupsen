package show

import (
	"errors"
	"log"
	"os/exec"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
show displays any tc settings jupsen has placed on a unit

Examples:

juju jupsen show wordpress/0
`

type Command struct {
	cmd.CommandBase
	unit string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "show",
		Args:    "unit",
		Purpose: "show displays any tc settings jupsen has placed on a unit",
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
	args = append(args, common.Show()...)
	log.Println(args)
	out, err := exec.Command("juju", args...).Output()
	if err != nil {
		return err
	}
	log.Println(string(out))
	return nil
}
