package main

import (
	"kanban/cmd"
	"kanban/data"
)

func main() {
	dummyData := data.Task{"12312312", "task 1", "todo"}
	dummyData1 := data.Task{"12312313", "task 2", "doing"}
	data.Tasks = append(data.Tasks, dummyData, dummyData1)
	cmd.Execute()
}
