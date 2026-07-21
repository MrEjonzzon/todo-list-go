package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  "Display all tasks with their status (Pending or Done).",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for i, task := range tasks {
			status := "Pending"
			if task.Done {
				status = "Done"
			}
			fmt.Printf("%d. %s [%s] %s %s %s %s\n", i+1, task.Description, status, "Created: ", task.CreatedAt, "Updated: ", task.UpdatedAt)
		}
	},
}
