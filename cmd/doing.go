package cmd

import (
	"fmt"
	"kanban/data"

	"github.com/spf13/cobra"
)

// doingCmd represents the doing command
var doingCmd = &cobra.Command{
	Use:   "doing",
	Short: "Move a task to DOING state.",
	Run: func(cmd *cobra.Command, arg []string) {
		// accepts only one argument - task id
		if len(arg) > 1 || len(arg) == 0 {
			fmt.Println("The doing command takes only one argumment i.e. task id.")
			return
		}
		task := data.FindTaskByID(arg[0])
		task.ChangeState("doing")
		fmt.Println("Task successfully moved to doing state.")
	},
}

func init() {
	rootCmd.AddCommand(doingCmd)
}
