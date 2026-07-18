package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  "Add a new task with a description to your task list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{Description: description, Done: false})
		saveTasks(tasks)
		fmt.Printf("Task added: %s\n", description)
	},
}

func loadTasks() []Task {
	file, err := os.Open("tasks.json")
	if os.IsNotExist(err) {
		return []Task{}
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []Task
	json.NewDecoder(file).Decode(&tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	file, err := os.Create("tasks.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(tasks)
}
