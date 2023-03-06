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


func submitHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	person := Person{Name: name, Email: email, Message: message}

	stmt, err := db.Prepare("INSERT INTO people(name, email, message) values(?,?,?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(person.Name, person.Email, person.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Thank you for your submission, %s!", person.Name)
}

/*The submitHandler function responds to requests to the /submit path when the form is submitted.
 This function adds a record to the SQLite database using the database/sql package. 
 First, the database connection is created using the sql.Open function, and then the database connection is closed by calling 
 db.Close() with the keyword defer.
 The db.Ping() function checks if the database connection is healthy. 
 If there is an error in the connection, an HTTP 500 error is generated with the http.Error function and the request is not responded to.
 The form data is parsed with the r.FormValue() function and a person variable is created using the structure named Person.
 The db.Prepare() function prepares an SQL statement in the database.*/




func viewHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT name, email, message FROM people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var name, email, message string
		if err := rows.Scan(&name, &email, &message); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		person := Person{Name: name, Email: email, Message: message}
		people = append(people, person)
	}

	tmpl, err := template.ParseFiles("view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, people)
}


/* Using the SQL query SELECT name, email, message FROM people, 
 this code selects data from the people table and converts that data into a slice of Person structures ([]Person).
 It then displays the data in an HTML page, using the view.html file as the HTML template.
 The SQL statement (SELECT name, email, message FROM people) used in this code selects all data. 
 In real-world scenarios, queries will often be more specific and the WHERE statement can be used to select, for example, only a specific user's data.


*/
