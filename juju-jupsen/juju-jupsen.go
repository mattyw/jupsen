package main

import (
	"fmt"
	"os"

	"github.com/juju/cmd"

	"github.com/mattyw/jupsen/fix"
	"github.com/mattyw/jupsen/flaky"
	"github.com/mattyw/jupsen/heal"
	"github.com/mattyw/jupsen/part"
	"github.com/mattyw/jupsen/show"
	"github.com/mattyw/jupsen/slow"
)

var commandDoc = `
Jupsen is used to cause network failures in juju environments.
All hail eris
`

func main() {
	ctx, err := cmd.DefaultContext()
	if err != nil {
		fmt.Printf("failed to get command context: %v\n", err)
		os.Exit(2)
	}
	jupsen := cmd.NewSuperCommand(cmd.SuperCommandParams{
		Name: "jupsen",
		Doc:  commandDoc,
		Log:  &cmd.Log{},
	})

	jupsen.Register(&part.Command{})
	jupsen.Register(&heal.Command{})
	jupsen.Register(&fix.Command{})
	jupsen.Register(&flaky.Command{})
	jupsen.Register(&slow.Command{})
	jupsen.Register(&show.Command{})

	args := os.Args
	os.Exit(cmd.Main(jupsen, ctx, args[1:]))
}
