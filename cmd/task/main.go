package main

import (
	//"github.com/yedeka/Go_Projects/cmd/task/cmd"
	"fmt"

	"github.com/yedeka/Go_Projects/cmd/task/config"
	"github.com/yedeka/Go_Projects/cmd/task/dao"
)

func main() {
	configurations, err := config.LoadConfigurations()
	if nil != err {
		fmt.Printf("Config error %v", err)
	}
	dbConfig := configurations.Database
	db, err := dao.ConnectToDb(dbConfig)
	if nil != err {
		fmt.Println("Error while connecting to Databse")
		fmt.Printf("%s", err.Error())
	}

	/*taskId, err := dao.CreateTask("Testing_Task", db)
	if nil != err {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("Created tasks successfully with Task Id %d", taskId)
	*/
	taskList, err := dao.ListAllTasks(db)
	if nil != err {
		fmt.Printf("%s", err.Error())
	}
	for _, task := range taskList {
		fmt.Printf("Task to be deleted => %+v \n", task)
		err := dao.DeleteTask(task.TaskId, db)

		if nil != err {
			fmt.Printf("Error while deleting the task %s", err.Error())
		}
	}

	defer db.TaskRepository.Close()
	//cmd.RootCmd.Execute()
}
