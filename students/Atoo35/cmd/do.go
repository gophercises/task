package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Atoo35/go-taskmanager-cli/db"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}

			task := tasks[id-1]
			// _, completedTaskErr := db.CreateCompletedTask(task.Name)
			// if completedTaskErr != nil {
			// 	fmt.Println("Something went wrong:", completedTaskErr.Error())
			// 	os.Exit(1)
			// }
			err := db.UpdateStatus(task.Key, "completed")
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
