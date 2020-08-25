package cmd

import (
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the TODO list",
	// 	Long: `A longer description that spans multiple lines and likely contains examples`

	Run: func(cmd *cobra.Command, args []string) {
		s := strings.Join(args, " ")
		id, err := db.AddTask(s)
		if err != nil {
			fmt.Printf("Unable to add task %s : %v\n", s, err)
		}
		fmt.Println("Added", s, "with id", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly, e.g.:
	// addCmd.Flags().StringVarP(&description, "desc", "d", "", "The task short description")
}
