package cmd

import (
	"fmt"
	"os"

	"github.com/Atoo35/go-taskmanager-cli/db"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Shows completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllCompletedTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No completed tasks to show!")
			return
		}
		fmt.Println("You have the following completed tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Name)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
