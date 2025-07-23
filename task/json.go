package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const path = "test.txt"

func JSONToFile(tracker TaskTracker) {

	jsonString := TaskTrackerToJSON(tracker)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error JSONtoFile:", err)
	}

	_, err = file.WriteString(jsonString)
	if err != nil {
		fmt.Println("Error JSONtoFile:", err)
	}
	defer func() {
		file.Close()
	}()
}

func FileToJSON() string {
	jsonString, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error FileToJSON:", err)
	}
	return string(jsonString)
}

func TaskTrackerToJSON(t TaskTracker) string {
	jsonString, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		fmt.Println("Error TaskTrackerToJSON:", err)
	}
	return string(jsonString)
}

func JSONToTaskTracker(jsonString string) TaskTracker {
	var tracker TaskTracker
	err := json.Unmarshal([]byte(jsonString), &tracker)
	if err != nil {
		fmt.Println("Error JSONToTaskTracker:", err)
	}
	return tracker
}
