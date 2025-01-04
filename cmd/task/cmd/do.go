package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			taskId, err := strconv.Atoi(arg)
			if nil != err {
				fmt.Printf("Error occured while parsing %s\n", arg)
			}
			ids = append(ids, taskId)
		}
		fmt.Printf("%v", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
