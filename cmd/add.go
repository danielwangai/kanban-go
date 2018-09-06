package cmd

import (
	"fmt"
	"kanban/data"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		if len(taskName) == 0 {
			fmt.Println("add called empty")
		}
		newTask := data.Task{Name: taskName}
		newTask.CreateTask()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
