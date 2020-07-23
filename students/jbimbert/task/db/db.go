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
		t := Task{Id: id, Desc: t, CreateTS: time.Now(), DoneTS: time.Now(), Status: 0}
		bs, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bs)
	})
	if e != nil {
		return -1, e
	}
	return id, nil
}

//ListAll list all tasks in the DB
func ListAll() []Task {
	var t []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		b.ForEach(func(k, v []byte) error {
			// id := btoi(k)
			task, err := Decode(v)
			if err != nil {
				return err
			}
			t = append(t, task)
			return nil
		})
		return nil
	})
	return t
}

// DoneTask set the status of the task with the given ID to DONE
// Do not change the ID of the task in the DB
func DoneTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		v := b.Get(itob(id))
		if v == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err := Decode(v)
		if err != nil {
			return err
		}
		(&t).Done()
		bytes, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bytes)
	})
}

// FindTask retrieve the task with the given ID
func FindTask(id int) (Task, error) {
	var t Task
	var err error
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		bytes := b.Get(itob(id))
		if bytes == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err = Decode(bytes)
		return err
	})
	return t, err
}

// DeleteTask remove a task from the DB
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
