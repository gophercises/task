package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

const formatTS = "2006-01-02 15:04:05" // display timestamp format

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display all tasks (use -a to display done/waived tasks, use -u/-c/-e to filter by urgency/criticality/effort)",

	Run: func(cmd *cobra.Command, args []string) {
		tasks := db.ListAll()
		if len(tasks) == 0 {
			fmt.Println("No tasks to view")
			return
		}
		if effort1 {
			sort.SliceStable(tasks, func(i, j int) bool {
				return tasks[i].Effor < tasks[j].Effor
			})
		}
		if criticality1 {
			sort.SliceStable(tasks, func(i, j int) bool {
				return tasks[i].Critic < tasks[j].Critic
			})
		}
		if urgency1 {
			sort.SliceStable(tasks, func(i, j int) bool {
				return tasks[i].Urge < tasks[j].Urge
			})
		}
		if all {
			fmt.Println("[TODO]")
		}
		for _, t := range tasks {
			if t.IsTodo() {
				fmt.Printf("%-3d %s %s %-7s %-10s  %s\n", t.Id, t.Urgency(), t.Criticality(), t.Effort(), t.CreateTS.Format(formatTS), t.Desc)
			}
		}
		if all {
			fmt.Println()
			fmt.Println("[DONE]")
			for _, t := range tasks {
				if t.IsDone() {
					fmt.Printf("%-3d %-10s  %-10s  %s\n", t.Id, t.CreateTS.Format(formatTS), t.DoneTS.Format(formatTS), t.Desc)
				}
			}
			fmt.Println()
			fmt.Println("[WAIVE]")
			for _, t := range tasks {
				if t.IsWaived() {
					fmt.Printf("%-3d %s %-10s  %-10s  %s\n", t.Id, t.CreateTS.Format(formatTS), t.DoneTS.Format(formatTS), t.State(), t.Desc)
				}
			}
		}
	},
}

var all bool                             // if false display only the TODO tasks, else display all tasks
var urgency1, criticality1, effort1 bool // filters

func init() {
	rootCmd.AddCommand(listCmd)
	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Display all tasks, done, give up, etc.")
	listCmd.Flags().BoolVarP(&urgency1, "urgency", "u", false, "Filter all tasks by urgency")
	listCmd.Flags().BoolVarP(&criticality1, "criticality", "c", false, "Filter all tasks by criticality")
	listCmd.Flags().BoolVarP(&effort1, "effort", "e", false, "Filter all tasks by effort")
}
