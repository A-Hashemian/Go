import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)
//These lines are importing the packages we will be using and the SQLite database driver.

type Person struct {
	Name    string
	Email   string
	Message string
}

//We define a structure named Person, where there are three properties named Name, Email and Message.
