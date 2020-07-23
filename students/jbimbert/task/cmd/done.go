package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Task is done",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				e := db.DoneTask(id)
				if e == nil {
					fmt.Printf("task with id %d done\n", id)
				} else {
					fmt.Printf("Error when doing task with id %d : %s\n", id, e)
				}
			} else {
				fmt.Println("Bad task id\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// doCmd.Flags().IntVarP(&taskId, "id", "i", -1, "The task ID (as displayed with list)")
}
