package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Veritabanı bağlantısı için değişkenler
// Variables for database connection
var db *sql.DB
var err error

// Kullanıcı verileri için bir yapı tanımlanmasi
// Defining a structure for user data
type User struct {
	ID       int
	Username string
	Password string
}

// HTTP isteklerini ele almak için bir yönlendirme mekanizması oluşumu
// Formation of a redirect mechanism to handle HTTP requests
func routes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

// HTTP GET isteğine yanıt vermek için bir işlev tanımlama
// Define a function to respond to an HTTP GET request
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Şablon dosyasını yüklenmesi
	// Loading the template file
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Şablon verileri için bir harita oluşturmasi
	// Generating a map for template data
	data := map[string]string{
		"Title": "Hoşgeldiniz",
		"Body":  "Bu bir test sayfasıdır.",
	}

// Şablonu veriyle doldurarak yanıt vermesi
// Reply by filling the template with data
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
// HTTP POST isteğine yanıt vermek için bir işlev tanımlama
// Define a function to respond to an HTTP POST request
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
