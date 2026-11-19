package main


import (
    "database/sql"
    "fmt"
	"os"
    _ "github.com/glebarez/go-sqlite"
)


func main() {
	homedir, _ := os.UserHomeDir()
	db, _ := sql.Open("sqlite", fmt.Sprintf("%s/.fasttrack/fasttrack.db", homedir))
}
