package bite

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func Bite() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bite",
		Short: "The critter bites",
		Long: `Run for your life. 
			When a critter bites it never stops`,
		Run: func(cmd *cobra.Command, args []string) {
			err := BiteNonStop(args)
			if err != nil {
				glog.Exitf("%s\n", err.Error())
			}
		},
	}
	cmd.Flags().IntVarP(&people, "people", "p", 1, "people to eat")
	return cmd
}
