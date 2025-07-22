package task

import "fmt"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UploatedAt  string `json:"uploatedAt"`
}

func (t *Task) New(a int, d string, s bool, c string, u string) {
	*t = Task{Id: a, Description: d, Status: s, CreatedAt: c, UploatedAt: u}
}

func (t Task) TaskPrint() {
	fmt.Printf("id: %d\n", t.Id)
	fmt.Printf("description: %s\n", t.Description)
	fmt.Printf("status: %t\n", t.Status)
	fmt.Printf("createdAt: %s\n", t.CreatedAt)
	fmt.Printf("uploatedAt: %s\n", t.UploatedAt)
}
