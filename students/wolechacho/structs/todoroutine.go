package structs

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

var (
	//Db is database Instance
	Db *bolt.DB
	er error
)

//Todo struct contains fields for todo items
type Todo struct {
	Item     string
	Isdone   bool
	ID       uint64
	Datetime time.Time
}

//ConnectDb opens Database for transactions
func ConnectDb() {
	Db, er = bolt.Open("task.db", 0600, nil)
	if er != nil {
		log.Fatal(er)
	}
}

//Itob takes an int64 to bytes
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//Gettodolist get all uncompleted task from bucket list
func Gettodolist() ([]Todo, error) {
	var err error
	todos := make([]Todo, 0)
	Db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		tx.CreateBucketIfNotExists([]byte("todos"))
		b := tx.Bucket([]byte("todos"))
		b.ForEach(func(k, v []byte) error {
			buffer := &bytes.Buffer{}
			buffer.Write(v)
			dec := gob.NewDecoder(buffer)

			t := &Todo{}
			err = dec.Decode(t)

			if err == nil {
				todos = append(todos, *t)
			} else {
				log.Fatal("decode error:", err)
			}
			return err
		})
		return err
	})
	return todos, err
}

//Savetodo add todo to the todos bucket list
func Savetodo(item string) error {
	return Db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("todos"))
		// Retrieve the users bucket.
		b := tx.Bucket([]byte("todos"))
		// Generate ID for the user.
		id, _ := b.NextSequence()
		t := Todo{
			Item:     item,
			Isdone:   false,
			ID:       id,
			Datetime: time.Now(),
		}
		buffer := &bytes.Buffer{}
		err := gob.NewEncoder(buffer).Encode(t)
		if err != nil {
			return err
		}
		// Persist bytes to users bucket.
		return b.Put(Itob(int(id)), buffer.Bytes())

	})
}

//Delete a todo item from the todos bucket list
func Delete(id uint64) error {
	var err error
	return Db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("todos"))
		b := tx.Bucket([]byte("todos"))
		err = b.Delete(Itob(int(id)))
		return err
	})
}

//Updatetodo update the todo information
func Updatetodo(key []byte) error {
	var err error
	return Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		val := b.Get([]byte(key))
		if val == nil {
			return errors.New("Key for the todo item not found")
		}
		buffer := &bytes.Buffer{}
		buffer.Write(val)
		decodedBytes := gob.NewDecoder(buffer)
		t := &Todo{}
		err = decodedBytes.Decode(t)
		if err != nil {
			return err
		}
		buffer.Reset()
		t.Isdone = true
		encodedBytes := gob.NewEncoder(buffer)
		err = encodedBytes.Encode(t)
		if err != nil {
			return err
		}
		return b.Put(key, buffer.Bytes())
	})
}
