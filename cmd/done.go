package cmd

import (
	"fmt"
	"kanban/data"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Move task from doing to done.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 || len(args) == 0 {
			fmt.Println("The done command takes only one argumment i.e. task id.")
			return
		}
		task := data.FindTaskByID(args[0])
		task.ChangeState("done")
		fmt.Println("Task successfully moved to done state.")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
