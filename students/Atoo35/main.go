package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Atoo35/go-taskmanager-cli/cmd"
	"github.com/Atoo35/go-taskmanager-cli/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
