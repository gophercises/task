package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Task is deleted from DB",

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
				fmt.Printf("Really delete task [%s]? [y/n]", t.Desc)
				fmt.Scanln(&input)
				if strings.Trim(input, " ") != "y" {
					return
				}
				err = db.DeleteTask(id)
				if err == nil {
					fmt.Printf("task with id %d deleted\n", id)
				} else {
					fmt.Printf("Error when removing task with id %d : %s\n", id, err)
				}
			} else {
				fmt.Println("Bad task id\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// doCmd.Flags().IntVarP(&taskId, "id", "i", -1, "The task ID (as displayed with list)")
}
