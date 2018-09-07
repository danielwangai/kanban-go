package data

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	Id    string
	Name  string
	State string // todo, doing, done
}

var Tasks []Task

func (task *Task) CreateTask() {
	/*
		creates a task
		use *task - to allow updating of the task
		use of task creates a new copy of itself hence update is not done
	*/
	if len(task.Name) == 0 {
		panic("Cannot create an empty task.")
	}
	for _, t := range Tasks {
		if t.Name == task.Name {
			panic("The task " + task.Name + " exists")
		}
	}
	task.Id = uuid.Must(uuid.NewV4()).String()
	task.State = "todo"
	Tasks = append(Tasks, *task)
	fmt.Println("Task created successfully.")
}

func (task *Task) ChangeState(state string) {
	/*
		state: the new desired state for the task
	*/
	if task.State == "todo" && state == "done" {
		panic("Cannot move task from todo to done directly.")
	}
	if task.State == state {
		panic("Cannot move a task to the same state.")
	}
	task.State = state
}

func FindTaskByID(id string) Task {
	//generic method to find Tasks by id
	var t Task
	for _, task := range Tasks {
		if task.Id == id {
			t = task
			break
		}
	}
	if t.Name == "" {
		panic("Task not found.")
	}
	return t
}
