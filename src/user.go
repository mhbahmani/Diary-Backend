package main

import (
	"net/http"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type User struct {
    gorm.Model
	Username  string
	Password string
	Email string
	FirstName string
	LastName string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    var users []User
    db.Find(&users)

    json.NewEncoder(w).Encode(users)
}