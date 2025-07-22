package task

import (
	"fmt"
	"strconv"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UploatedAt  string `json:"uploatedAt"`
}

func (t *Task) New(a int, d string, s int, c string, u string) {
	*t = Task{Id: a, Description: d, Status: s, CreatedAt: c, UploatedAt: u}
}

func (t Task) TaskPrint() {
	fmt.Printf("id: %d\n", t.Id)
	fmt.Printf("description: %s\n", t.Description)
	fmt.Printf("status: %s\n", t.GetStatus())
	fmt.Printf("createdAt: %s\n", t.CreatedAt)
	fmt.Printf("uploatedAt: %s\n\n", t.UploatedAt)
}

func (t Task) GetStatus() string {
	switch t.Status {
	case 0:
		return "todo"
	case 1:
		return "in progress"
	case 2:
		return "done"
	default:
		return "Error status..."
	}
}

type TaskTracker struct {
	Size      int
	TaskTrack []Task
}

func (t *TaskTracker) Add(str string) {
	if t.FindTaskInTracker(str) {
		fmt.Printf("Error, this Task alrady exist\n\n")
		return
	}
	task := t.SetTrack(str)
	t.Size++
	t.TaskTrack = append(t.TaskTrack, task)
	fmt.Printf("Task added successfully (ID: %d)\n\n", t.Size)
}

func (t *TaskTracker) SetTrack(str string) Task {
	id := t.Size + 1
	status := false
	year := strconv.Itoa(time.Now().Year())
	mo := int(time.Now().Month())
	month := ""
	if mo < 10 {
		month = "0" + strconv.Itoa(mo)
	} else {
		month = strconv.Itoa(mo)
	}
	d := time.Now().Day()
	day := ""
	if d < 10 {
		day = "0" + strconv.Itoa(d)
	} else {
		day = strconv.Itoa(d)
	}
	h := time.Now().Hour()
	hours := ""
	if h < 10 {
		hours = "0" + strconv.Itoa(h)
	} else {
		hours = strconv.Itoa(h)
	}
	m := time.Now().Minute()
	minute := ""
	if m < 10 {
		minute = "0" + strconv.Itoa(m)
	} else {
		minute = strconv.Itoa(m)
	}
	createdAt := day + "." + month + "." + year + ", " + hours + ":" + minute
	uploatedAt := ""
	return Task{id, str, status, createdAt, uploatedAt}
}

func (t *TaskTracker) FindTaskInTracker(str string) bool {
	for i := 0; i < t.Size; i++ {
		if t.TaskTrack[i].Description == str {
			return true
		}
	}
	return false
}

func (t TaskTracker) Print() {
	for i := 0; i < t.Size; i++ {
		t.TaskTrack[i].TaskPrint()
	}
}

func (t *TaskTracker) Update(num int, str string) {
	if num > t.Size {
		fmt.Printf("Error, this ID (%d) is not found...\n\n", num)
		return
	}
	t.TaskTrack[num-1].Description = str
	fmt.Printf("Task %d is updated\n\n", num)
}

func (t *TaskTracker) Delete(num int) {
	if num > t.Size {
		fmt.Printf("Error, this ID (%d) is not found...\n\n", num)
		return
	}
	copy(t.TaskTrack[num-1:], t.TaskTrack[num:])
	t.TaskTrack = t.TaskTrack[:t.Size-1]
	t.Size--
	for i := num - 1; i < t.Size; i++ {
		t.TaskTrack[i].Id = i + 1
	}
	fmt.Printf("Task %d is deleted\n\n", num)
}
