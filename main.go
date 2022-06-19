package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

type Auth struct {
	gorm.Model
	ID        string
	authLevel authLevel
}
type authLevel uint8

const (
	ERROR  = 0
	LOW    = 1
	MEDIUM = 2
	HIGH   = 3
)

func createDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func messageWriter(writer http.ResponseWriter, message string) {
	fmt.Fprintf(writer, message)
}
func checkID(requestID string, db *gorm.DB) authLevel {
	var level authLevel
	level = ERROR
	query := db.Where("id = ? ", requestID).Select("authLevel")

	return level
}
func main() {
	db := createDatabase()
	sub := "/check-id/"
	http.HandleFunc(sub, func(writer http.ResponseWriter, request *http.Request) {
		id := request.RequestURI[len(sub):]
		message := "Requested auth for ID " + id
		messageWriter(writer, message)
		checkID(id, db)
	})
	http.ListenAndServe("localhost:8080", nil)
}
