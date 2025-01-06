package dao

import (
	"encoding/binary"
	"fmt"
	"log"
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
func CreateTask(task string, db *TaskDb) (int, error) {
	var id int
	err := db.TaskRepository.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(db.BucketName)
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
func ListAllTasks(db *TaskDb) ([]Task, error) {
	var tasklist []Task
	err := db.TaskRepository.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(db.BucketName)
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
