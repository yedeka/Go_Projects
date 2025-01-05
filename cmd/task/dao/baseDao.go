package dao

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/yedeka/Go_Projects/cmd/task/config"
)

type TaskDb struct {
	TaskRepository *bolt.DB
}

func ConnectToDb(dbConfig config.DatabaseConfiguration) (*TaskDb, error) {
	taskDb := &TaskDb{}
	taskRepository, err := bolt.Open(dbConfig.DbName,
		0600,
		&bolt.Options{Timeout: time.Duration(dbConfig.DbTimeout) * time.Second})

	if nil != err {
		log.Fatal(err)
		return nil, fmt.Errorf("error while connecting to Task Database")
	}
	taskDb.TaskRepository = taskRepository
	fmt.Printf("Bucket name passed from config => %s", dbConfig.DbBucketName)
	// Create Task bucket if it does not already exist
	repositoryGenerator := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbConfig.DbBucketName))
		return err
	}
	taskRepository.Update(repositoryGenerator)
	return taskDb, nil
}
