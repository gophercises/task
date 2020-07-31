package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Use this command to finish a task in the database",
	Long:  "",
	Run:   doTasks,
}

func init() {

}

func doTasks(cmd *cobra.Command, args []string) {
	fmt.Println("haha that does a list")
}
