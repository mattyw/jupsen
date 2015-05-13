package part

import (
	"errors"
	"log"
	"os/exec"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/common"
)

var commandDoc = `
part causes network partiions between the two units

Examples:

juju jupsen part wordpess/0 mysql/0 

The partition will occur on the machine, so any colocated units will be affected
`

type Command struct {
	cmd.CommandBase
	unitA string
	unitB string
}

func (c *Command) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "part",
		Args:    "unitA unitB",
		Purpose: "partitions unitA and unitB",
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
	args = append(args, common.Part(unitBIP)...)
	log.Println(args)
	_, err = exec.Command("juju", args...).Output()
	if err != nil {
		return err
	}
	return nil
}
