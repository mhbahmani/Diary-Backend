package main

import (
	"fmt"

	"net/http"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
    gorm.Model
	Username  string
	Password string
	Email string
	FirstName string
	LastName string
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()
	defer db.Close()
    var users []User
    db.Find(&users)

    json.NewEncoder(w).Encode(users)
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	db := connectToDB()
	defer db.Close()

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)
	fmt.Fprintf(w, "ok") // User successfully created
}