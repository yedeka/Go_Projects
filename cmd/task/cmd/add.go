package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run: func(cmd *cobra.Command, arguments []string) {
		fmt.Println("Add called")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
