package main

import (
	"kanban/cmd"
	"kanban/data"
)

func main() {
	dummyData := data.Task{"12312312", "task 1", "todo"}
	data.Tasks = append(data.Tasks, dummyData)
	cmd.Execute()
}
