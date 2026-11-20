package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".fasttrack")
	os.MkdirAll(dir, os.ModePerm)

	dsn := filepath.Join(dir, "fasttrack.db") +
		"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"
	db, err := sql.Open("sqlite", dsn)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlsetup := "CREATE TABLE IF NOT EXISTS \"projects\" (\n\t\"id\" TEXT NOT NULL UNIQUE,\n\t\"name\" VARCHAR,\n\t\"directory\" VARCHAR,\n\t\"git-local\" VARCHAR,\n\t\"git-online\" VARCHAR,\n\t\"created\" TIMESTAMP,\n\t\"modified\" TIMESTAMP,\n\tPRIMARY KEY(\"id\"),\n\tFOREIGN KEY (\"id\") REFERENCES \"sessions\"(\"project\")\n\tON UPDATE NO ACTION ON DELETE NO ACTION\n);\n\nCREATE TABLE IF NOT EXISTS \"t1-heartbeats\" (\n\t\"id\" TEXT NOT NULL,\n\t\"session\" TEXT,\n\t\"relative-path\" VARCHAR,\n\t\"timestamp\" TIMESTAMP,\n\tPRIMARY KEY(\"id\"),\n\tFOREIGN KEY (\"session\") REFERENCES \"sessions\"(\"id\")\n\tON UPDATE NO ACTION ON DELETE NO ACTION\n);\n\nCREATE TABLE IF NOT EXISTS \"sessions\" (\n\t\"id\" TEXT NOT NULL,\n\t\"started\" TIMESTAMP,\n\t\"ended\" TIMESTAMP,\n\t\"project\" TEXT,\n\tPRIMARY KEY(\"id\")\n);\n"

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
		db.Exec(sqlsetup)
	}

}
