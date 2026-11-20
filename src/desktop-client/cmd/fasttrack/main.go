package main

import (
	"fmt"
	"os"

	_ "github.com/glebarez/go-sqlite"
	"hexagonical.ch/timetrack-desktop-client/internal/store"
)

func main() {
	fmt.Println("Hello, Hackclub")
	if len(os.Args) < 2 {
		fmt.Println("Usage: fasttrack <command>")
		return
	}
	args := os.Args[1:]
	pwd, _ := os.Getwd()
	switch args[0] {
	case "project":
		if len(args) < 2 {
			fmt.Println("Usage: fasttrack project <init>")
			return
		}
		switch args[1] {
		case "init":
			db, err := store.Open(store.GetDefaultPath())
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to open database: %v\n", err)
				return
			}
			defer db.Close()

			project, err := db.GetProjectByDir(pwd)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to look up project: %v\n", err)
				return
			}
			if project != nil {
				fmt.Println("Project already exists in this location?")
				return
			}

			name, ok := promptProjectName()
			if !ok {
				fmt.Println("Project init cancelled.")
				return
			}
			project, _ = db.CreateProject(name, pwd)
			fmt.Printf("Project %q initialized in %s\n", name, pwd)

		}
	case "sync":
		fmt.Println("sync")
	}

}
