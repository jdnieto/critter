package cmd

import (
	"github.com/jdnieto/critter/bite"
	"github.com/jdnieto/critter/roll"
	"github.com/spf13/cobra"
)

func Action() *cobra.Command {
	action := &cobra.Command{
		Use:   "action",
		Short: "Actions allowed bite/roll",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
	action.AddCommand(bite.Bite(), roll.Roll())
	return action
}
