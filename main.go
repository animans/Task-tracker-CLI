package main

import (
	"Task-tracker-CLI/task"
	"fmt"
	"os"
	"strconv"
)

func main() {
	command := os.Args[1]
	tracker := task.TaskTracker{}
	if task.FileExists() {
		tracker = task.JSONToTaskTracker(task.FileToJSON())
	}
	arg := func() int {
		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		return arg
	}
	switch command {
	case "add":
		tracker.Add(os.Args[2])
	case "update":
		tracker.Update(arg(), os.Args[3])
	case "delete":
		tracker.Delete(arg())
	case "mark-in-progress":
		tracker.MarkProgress(arg(), 1)
	case "mark-done":
		tracker.MarkProgress(arg(), 2)
	case "list":
		if len(os.Args) == 3 {
			tracker.PrintArg(os.Args[2])
		} else if len(os.Args) == 2 {
			tracker.Print()
		} else {
			fmt.Println("Error command")
		}

	default:
		fmt.Println("Error command")
	}
	task.JSONToFile(tracker)
}
