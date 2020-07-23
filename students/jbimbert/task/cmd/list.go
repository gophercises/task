package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"

	"github.com/spf13/cobra"
)

const formatTS = "2006-01-02 15:04:05" // display timestamp format

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display the list of tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ListAll()
		if len(tasks) == 0 {
			fmt.Println("No tasks to view")
			return
		}
		if all {
			fmt.Println("[TODO]")
		}
		for _, t := range tasks {
			if t.IsTodo() {
				fmt.Printf("%-3d %-10s %s\n", t.Id, t.CreateTS.Format(formatTS), t.Desc)
			}
		}
		if all {
			fmt.Println()
			fmt.Println("[DONE]")
			for _, t := range tasks {
				if t.IsDone() {
					fmt.Printf("%-3d %-10s %-10s %s\n", t.Id, t.CreateTS.Format(formatTS), t.DoneTS.Format(formatTS), t.Desc)
				}
			}
			fmt.Println()
			fmt.Println("[GIVE UP]")
			for _, t := range tasks {
				if t.IsGiveUp() {
					fmt.Printf("%-3d %s %-10s %-10s %s\n", t.Id, t.CreateTS.Format(formatTS), t.DoneTS.Format(formatTS), t.State(), t.Desc)
				}
			}
		}
	},
}

var all bool // if false display only the TODO tasks, else display all tasks

func init() {
	rootCmd.AddCommand(listCmd)
	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Display all tasks, done, give up, etc.")
}
