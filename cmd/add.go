package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// Source - https://stackoverflow.com/a/64632817
// Posted by colm.anseo
// Retrieved 2026-07-18, License - CC BY-SA 4.0

type autoInc struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         int
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}

var ai autoInc // global instance

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var addCmd = &cobra.Command{
	Use:   "add [task description]",
	Short: "Add a new task",
	Long:  "Add a new task with a description to your task list.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{
			Id:          ai.ID(),
			Description: description,
			Done:        false,
			CreatedAt:   time.Now().Format("2006-01-02T15:04:05Z"),
			UpdatedAt:   time.Now().Format("2006-01-02T15:04:05Z"),
		})
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
