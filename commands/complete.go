package commands

import (
	"fmt"
	"strconv"

	"todo-app/task"
)

func CompleteTask(idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID:", idStr)
		return
	}

	tasks, err := task.ReadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			err = task.WriteTasks("tasks.json", tasks)
			if err != nil {
				fmt.Println("Error writing tasks:", err)
				return
			}
			fmt.Println("Task completed successfully!")
			return
		}
	}

	fmt.Println("Task not found:", id)
}
