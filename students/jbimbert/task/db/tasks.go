package db

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Task represents a row in the DB
type Task struct {
	Id       int       // task id
	Desc     string    // task description
	CreateTS time.Time // when the task was created
	DoneTS   time.Time // when the task was done
	Status   int       // 0 todo, 1 doing, 2 done, 3 waive
}

// Status of the task
func (t Task) State() string {
	switch t.Status {
	default:
		return "TODO"
	case 1:
		return "DOING"
	case 2:
		return "DONE"
	case 3:
		return "WAIVE"
	}
}

// Update task description with the given string
func (t *Task) UpdateDesc(s string) {
	t.Desc = s
}

// Todo set task status to TODO
func (t *Task) Todo() {
	t.Status = 0
	t.DoneTS = time.Now()
}

// Done set task status to DONE
func (t *Task) Done() {
	t.Status = 2
	t.DoneTS = time.Now()
}

// Done set task status to WAIVW
func (t *Task) Waive() {
	t.Status = 3
	t.DoneTS = time.Now()
}

// IsTodo true is the task status is TODO
func (t Task) IsTodo() bool {
	return t.Status == 0
}

// IsDone true is the task status is DONE
func (t Task) IsDone() bool {
	return t.Status == 2
}

// IsWaived true is the task status is WAIVE
func (t Task) IsWaived() bool {
	return t.Status == 3
}

// Serialize a task into an array of bytes
func (t Task) toByte() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode deserialize an array of bytes into a task
func Decode(b []byte) (Task, error) {
	pbuf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(pbuf)
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}
