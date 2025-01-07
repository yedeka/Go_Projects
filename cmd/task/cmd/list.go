package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yedeka/Go_Projects/cmd/task/dao"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Provides a list of all the pending tasks",
	Run: func(cmd *cobra.Command, args []string) {
		taskList, err := dao.ListAllTasks()
		if nil != err {
			fmt.Printf("%s", err.Error())
		}
		if len(taskList) == 0 {
			fmt.Println("You have no pending tasks under your name. Feel free to add new tasks")
		} else {
			fmt.Print("You have the following tasks: \n")
			for _, task := range taskList {
				fmt.Printf("%d. %s \n", task.TaskId, task.TaskName)
			}
			fmt.Println("")
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
