package main

import (
    "fmt"

	"net/http"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func connectToDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../sqlite.db")
    if err != nil {
        panic("failed to connect database")
    }
	return db
}

func initialMigration() {
	db := connectToDB()
	defer db.Close()

    db.AutoMigrate(&User{})
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Welcome! This is home page.</h1>")
}

func handleRequests(){
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/users", getAllUsers) // users
	http.HandleFunc("/register.php", createNewUser) // users

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {	
	initialMigration()

	handleRequests()
}
