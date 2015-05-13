package heal

import (
	"errors"
	"log"
	"os/exec"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
heal heals any network partitions caused by the part command

Examples:

juju jupsen heal wordpress/0 mysql/0

`

type Command struct {
	cmd.CommandBase
	unitA string
	unitB string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "heal",
		Args:    "unitA unitB",
		Purpose: "heals partitions between unitA and unitB",
		Doc:     commandDoc,
	}
}

func (c *Command) Init(args []string) error {
	if len(args) < 2 {
		return errors.New("expected two unit names")
	}
	c.unitA = args[0]
	c.unitB = args[1]
	return cmd.CheckEmpty(args[2:])
}

func (c *Command) Run(ctx *cmd.Context) error {
	unitBIP, err := common.UnitPrivateIp(c.unitB)
	if err != nil {
		return err
	}
	args := []string{"ssh", c.unitA}
	args = append(args, common.Heal(unitBIP)...)
	log.Println(args)
	_, err = exec.Command("juju", args...).Output()
	if err != nil {
		return err
	}
	return nil
}
