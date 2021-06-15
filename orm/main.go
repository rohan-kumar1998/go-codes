package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setup(db *gorm.DB) {
	db.AutoMigrate(&Syllabus{}, &Concept{}, &ConceptQuestion{}, &Journey{}, &Node{}, &Question{}, &Subject{}, &Topic{}, &Video{}, &VideoSegment{})
}

var Db *gorm.DB

func main() {
	dsn := "root:welcome123@tcp(localhost:3306)/db_bfs_toppr"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error in sql.open")
		panic(err.Error())
	}
	setup(Db)
	PreloadSyllabus(Db)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/api/v1/syllabus/", GetSyllabus)
	http.HandleFunc("/api/v1/syllabus/subjects/", GetSubjects)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Successful Connection to Database!")
}
