package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run: func(cmd *cobra.Command, arguments []string) {
		taskName := strings.Join(arguments, " ")
		fmt.Printf("Added \"%s\" to your list", taskName)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
