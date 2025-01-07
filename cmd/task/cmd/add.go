package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yedeka/Go_Projects/cmd/task/dao"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run: func(cmd *cobra.Command, arguments []string) {
		taskName := strings.Join(arguments, " ")
		_, err := dao.CreateTask(taskName)
		if nil != err {
			fmt.Printf("error while adding your task %s", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" to your list\n", taskName)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
