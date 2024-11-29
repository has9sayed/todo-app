package todo

import "fmt"

type Task struct {
	ID int
	Title string
	Done bool
}

type TaskList struct {
	Tasks []Task
	NextID int
}

func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: []Task{},
		NextID: 1,
	}
}

func (tl *TaskList) AddTask(title string) {
	task := Task {
		ID: tl.NextID,
		Title : title,
		Done: false,
	}
	tl.Tasks = append(tl.Tasks, task)
	tl.NextID++
}

func (tl *TaskList) ListTasks() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tl.Tasks {
		status := "X"
		if task.Done {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n",status, task.ID, task.Title)
	}
}

func (tl *TaskList) MarkDone(id int) {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Done = true
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

func (tl *TaskList) DeleteTask(id int) {
	for i, task := range tl.Tasks {
		if task.ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			fmt.Printf("Task %d deleted.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
}