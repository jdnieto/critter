package bite

import (
	"fmt"
	"github.com/golang/glog"
	"strings"
)

var people int

func BiteNonStop(args []string) error {
	glog.V(2).Infof("Eating %s", args)
	for i := 0; i < people; i++ {
		fmt.Println(strings.Join(args, " "))
	}
	return nil
}
