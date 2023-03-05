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
