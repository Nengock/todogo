package commands

import (
	"fmt"

	"todo-app/task"
)

func ListTasks() {
	tasks, err := task.ReadTasks("tasks.json")
	if err != nil {
		fmt.Println("Error reading tasks:", err)
		return
	}

	for _, t := range tasks {
		status := "Incomplete"
		if t.Completed {
			status = "Complete"
		}
		fmt.Printf("%d: %s [%s]\n", t.ID, t.Description, status)
	}
}
