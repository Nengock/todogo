package commands

import (
	"fmt"

	"todo-app/task"
)

func AddTask(description string) {
	tasks, err := task.ReadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := task.Task{
		ID:          newID,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	err = task.WriteTasks("tasks.json", tasks)
	if err != nil {
		fmt.Println("Error writing tasks:", err)
		return
	}

	fmt.Println("Task added successfully!")
}
