/*
Package cmd Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"task/structs"
	"time"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your completed tasks",
	Long:  "List all of your completed tasks for the day",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := structs.Gettodolist()
		if err != nil {
			fmt.Printf("Error Retrieving Task : %s", err)
		} else {
			if len(todos) == 0 {
				fmt.Println("No task has been completed")
			}
			fmt.Println("You have the following tasks :")
			for _, t := range todos {
				if t.Datetime.Day() == time.Now().Day() {
					fmt.Printf("%d.  %s\n", t.ID, t.Item)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
