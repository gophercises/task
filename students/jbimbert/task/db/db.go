package db

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const (
	TasksTbl = "TASKTBL"
)

type Task struct {
	Id   int
	Desc string
}

var db *bolt.DB

func InitDb(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		return err
	}
	return createTable(TasksTbl)
}

func CloseDb() {
	db.Close()
}

// Create a table if it does not already exist
func createTable(tablename string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(tablename))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err.Error())
		}
		return nil
	})
}

// AddTask add a task "t" to the current table of tasks
// return its ID
func AddTask(t string) (int, error) {
	var id int
	e := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		uid, _ := b.NextSequence()
		id = int(uid)
		return b.Put(itob(id), []byte(t))
	})
	if e != nil {
		return -1, e
	}
	return id, nil
}

func ListAll() []Task {
	var t []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		b.ForEach(func(k, v []byte) error {
			id := btoi(k)
			t = append(t, Task{Id: id, Desc: string(v)})
			return nil
		})
		return nil
	})
	return t
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		c := b.Cursor()
		k, _ := c.Seek(itob(id))
		if k == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		c.Delete()
		return nil
	})
}

// itob returns an 8-byte big endian representation of v.
// used to create keys
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns the int representation of the 8-byte big endian b
// used to decode keys
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
