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

//defining a structure named Person, where there are three properties named Name, Email and Message.

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

//The formHandler function responds to requests to the / path and displays the HTML form.
