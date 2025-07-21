package task

import "fmt"

type Task struct {
	id          int
	description string
	status      bool
	createdAt   string
	uploatedAt  string
}

func (t *Task) New(a int, d string, s bool, c string, u string) {
	task := Task{id: a, description: d, status: s, createdAt: c, uploatedAt: u}
	*t = task
}

func (t Task) TaskPrint() {
	fmt.Printf("id: %d\n", t.id)
	fmt.Printf("description: %s\n", t.description)
	fmt.Printf("status: %t\n", t.status)
	fmt.Printf("createdAt: %s\n", t.createdAt)
	fmt.Printf("uploatedAt: %s\n", t.uploatedAt)
}
