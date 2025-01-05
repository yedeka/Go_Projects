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
	db, err := dao.ConnectToDb(configurations.Database)
	if nil != err {
		fmt.Printf("%s", err.Error())
	}
	defer db.TaskRepository.Close()
	fmt.Println("Finished creating db and Tasks bucket")
	//cmd.RootCmd.Execute()
}
