package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "use this command to add a task in the database",
	Long:  "",
	Run:   addTasks,
}

func init() {
}

func addTasks(cmd *cobra.Command, args []string) {
	fmt.Println("haha that adding a list")
}
