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

func (task *Task) changeState(state string) {
	/*
		state: the new desired state for the task
	*/
	if task.state == "todo" && state == "done" {
		panic("Cannot move task from todo to done directly.")
	}
	if task.state == state {
		panic("Cannot move a task to the same state.")
	}
	task.state = state
}

func findTaskByID(id string) Task {
	//generic method to find tasks by id
	var t Task
	for _, task := range tasks {
		if task.id == id {
			t = task
			break
		}
	}
	if t.name == "" {
		panic("Task not found.")
	}
	return t
}

func main() {
	x := Task{name: "one"}
	y := Task{name: "Two"}
	x.CreateTask()
	y.CreateTask()
	y.changeState("doing")
	fmt.Println(y)
}
