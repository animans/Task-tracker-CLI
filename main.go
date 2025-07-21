package main

import (
	"Task-tracker-CLI/task"
	"fmt"
)

func main() {
	//task := task.Task{id:18, "age", true, "13.03.2025, 23:25", "14.03.2025, 20:25"}
	var task task.Task
	task.New(18, "age", true, "13.03.2025, 23:25", "14.03.2025, 20:25")
	task.TaskPrint()
	fmt.Println()
}
