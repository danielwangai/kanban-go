package main

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Task struct {
	id    string
	name  string
	state string // todo, doing, done
}

var tasks []Task

func (task *Task) CreateTask() {
	/*
		creates a task
		use *task - to allow updating of the task
		use of task creates a new copy of itself hence update is not done
	*/
	if len(task.name) == 0 {
		panic("Cannot create an empty task.")
	}
	for _, t := range tasks {
		if t.name == task.name {
			panic("The task " + task.name + " exists")
		}
	}
	task.id = uuid.Must(uuid.NewV4()).String()
	task.state = "todo"
	tasks = append(tasks, *task)
	fmt.Println("Task created successfully.")
}

func main() {
	x := Task{name: "one"}
	y := Task{name: "Two"}
	x.CreateTask()
	y.CreateTask()
}
