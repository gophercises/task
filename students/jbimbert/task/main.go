package main

import (
	"Gophercizes/task/students/jbimbert/task/cmd"
	"Gophercizes/task/students/jbimbert/task/db"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	h, err := homedir.Dir()
	must(err)
	home := filepath.Join(h, "tasks.db")
	must(db.InitDb(home))
	defer db.CloseDb()

	cmd.Execute()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
