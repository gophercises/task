package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the task description",

	Run: func(cmd *cobra.Command, args []string) {
		for _, a := range args {
			id, err := strconv.Atoi(a)
			if err == nil {
				t, err := db.FindTask(id)
				if err != nil {
					fmt.Println(err)
					return
				}
				var input string
				fmt.Println("Enter new description for :", t.Desc)
				fmt.Scanf("%q", &input)
				if strings.Trim(input, " ") == "" {
					return
				}
				e := db.UpdateTask(id, input)
				if e == nil {
					fmt.Printf("task with id %d updated\n", id)
				} else {
					fmt.Printf("Error when updating task with id %d : %s\n", id, e)
				}
			} else {
				fmt.Println("Bad task id\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// doCmd.Flags().IntVarP(&taskId, "id", "i", -1, "The task ID (as displayed with list)")
}
