package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

// type Task struct {
// 	Key   int
// 	Value string
// }

type Task struct {
	Key    int    `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var newTask Task = Task{Name: task, Status: "pending"}
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(int(id64))
		marshlled, err := json.Marshal(newTask)
		if err != nil {
			return err
		}
		return b.Put(key, marshlled)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			if task.Status == "pending" {
				tasks = append(tasks, Task{
					Key:    btoi(k),
					Name:   task.Name,
					Status: task.Status,
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func AllCompletedTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				return err
			}
			if task.Status == "completed" {
				tasks = append(tasks, Task{
					Key:    btoi(k),
					Name:   task.Name,
					Status: task.Status,
				})
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func UpdateStatus(key int, status string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		var task Task
		json.Unmarshal(b.Get(itob(key)), &task)
		task.Status = status
		marshlled, err := json.Marshal(task)
		if err != nil {
			return err
		}
		return b.Put(itob(key), marshlled)
	})
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
