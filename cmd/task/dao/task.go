package dao

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/yedeka/Go_Projects/cmd/task/config"
)

type TaskDb struct {
	TaskRepository *bolt.DB
	BucketName     []byte
}

type Task struct {
	TaskName string
	TaskId   int
}

var taskDbConnection *TaskDb

func init() {
	configurations, err := config.LoadConfigurations()
	if nil != err {
		fmt.Printf("Config error %v", err)
		os.Exit(3)
	}
	dbConfig := configurations.Database
	db, err := ConnectToDb(dbConfig)
	if nil != err {
		fmt.Println("Error while connecting to Databse")
		fmt.Printf("%s", err.Error())
	}
	taskDbConnection = db
}

// ConnectToDb creates a connection to BoltDB and returns the pointer to corresponding connection or an error if any error is encoutered
// during connection creation.
func ConnectToDb(dbConfig config.DatabaseConfiguration) (*TaskDb, error) {
	taskDb := &TaskDb{BucketName: []byte(dbConfig.DbBucketName)}
	taskRepository, err := bolt.Open(dbConfig.DbName,
		0600,
		&bolt.Options{Timeout: time.Duration(dbConfig.DbTimeout) * time.Second})

	if nil != err {
		log.Fatal(err)
		return nil, fmt.Errorf("error while connecting to Task Database")
	}
	taskDb.TaskRepository = taskRepository
	// Create Task bucket if it does not already exist
	repositoryGenerator := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbConfig.DbBucketName))
		return err
	}
	err = taskRepository.Update(repositoryGenerator)
	if nil != err {
		return nil, fmt.Errorf("error while creating Task bucket")
	}
	return taskDb, nil
}

// CreateTask takes in the bucketName and Task name alongwith connection object to create the Connection and returns the task ID or error
// If any error is encountered while storing the task in Db.
func CreateTask(task string) (int, error) {
	var id int
	err := taskDbConnection.TaskRepository.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskDbConnection.BucketName)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))
	})

	if nil != err {
		return -1, err
	}

	return id, nil
}

// ListAllTasks -Lists all the tasks added to db in the bucket for the TODO application.
func ListAllTasks() ([]Task, error) {
	var tasklist []Task
	err := taskDbConnection.TaskRepository.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskDbConnection.BucketName)
		cursor := bucket.Cursor()

		for key, value := cursor.First(); nil != key; key, value = cursor.Next() {
			tasklist = append(tasklist, Task{
				TaskId:   btoi(key),
				TaskName: string(value),
			})
		}
		return nil
	})

	if nil != err {
		return nil, err
	}
	return tasklist, nil
}

// deleteTask takes in a taskId for the task to be deleted and deletes the task for given Id
func DeleteTask(taskId int) error {
	return taskDbConnection.TaskRepository.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskDbConnection.BucketName)
		return bucket.Delete(itob(taskId))
	})
}

// itob function takes an integer and gives a byte array used for bucket key
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoI takes a byte array and gives an integer to be used for converting db keys to displayable integers.
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
