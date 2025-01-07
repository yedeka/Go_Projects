package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yedeka/Go_Projects/cmd/task/dao"
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
			} else {
				ids = append(ids, taskId)
			}
		}

		taskList, err := dao.ListAllTasks()
		if nil != err {
			fmt.Printf("%s", err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(taskList) {
				fmt.Println("Invalid Task Number")
				continue
			}
			task := taskList[id-1]
			err := dao.DeleteTask(task.TaskId)

			if nil != err {
				fmt.Printf("Failed to mark %d task as complete. Error %s\n", id, err)
			} else {
				fmt.Printf("Marked %d task as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
