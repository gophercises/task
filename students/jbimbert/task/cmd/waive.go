package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// waiveCmd represents the waive command
var waiveCmd = &cobra.Command{
	Use:   "waive",
	Short: "Abandon or postpon a task (it is not deleted and you can reclaim it with -r)",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				e := db.WaiveTask(id, reclaim)
				if e == nil {
					if reclaim {
						fmt.Printf("task with id %d is reclaimed\n", id)
					} else {
						fmt.Printf("task with id %d is waived\n", id)
					}
				} else {
					fmt.Printf("Error when processing task with id %d : %s\n", id, e)
				}
			} else {
				fmt.Println("Bad task id\n", id)
			}
		}
	},
}

var reclaim bool

func init() {
	rootCmd.AddCommand(waiveCmd)

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	waiveCmd.Flags().BoolVarP(&reclaim, "reclaim", "r", false, "Reclaim the task which is already waived")
}
