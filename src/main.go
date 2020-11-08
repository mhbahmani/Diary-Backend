package main

import (
    "fmt"

	"net/http"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initialMigration() {
    db, err := gorm.Open("sqlite3", "../sqlite.db")
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&User{})
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Welcome! This is home page.</h1>")
}

func handleRequests(){
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/users", allUsers)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {	
	initialMigration()

	handleRequests()
}
