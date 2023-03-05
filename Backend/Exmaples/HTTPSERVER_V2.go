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
