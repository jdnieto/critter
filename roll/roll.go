package roll

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func Roll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roll",
		Short: "The critter rolls",
		Long: `Run for your life. 
			When a critter rolls it never stops`,
		Run: func(cmd *cobra.Command, args []string) {
			glog.V(2).Info("Rolling")
			fmt.Println("Rolling non stop")
		},
	}
	return cmd
}
