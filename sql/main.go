package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type data struct {
	SylName          string `json:"sylName"`
	SubjectName      string `json:"subjectName"`
	TopicName        string `json:"topicName"`
	SubtopicName     string `json:"subtopicName"`
	ConceptName      string `json:"conceptName"`
	VideoSegmentName string `json:"videosegmentName"`
	VideoName        string `json:"videoName"`
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Success")
}

func getSyllabus(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		req.ParseForm()

		id, err := strconv.Atoi(req.FormValue("id"))
		check(err)

		q, err := ioutil.ReadFile("./queries/syl-topics-vids.sql")
		check(err)

		rows, err := db.Query(string(q), id)
		check(err)
		var currentData data
		var allData []data = []data{}
		for rows.Next() {
			err = rows.Scan(&currentData.SylName,
				&currentData.SubjectName,
				&currentData.TopicName,
				&currentData.SubtopicName,
				&currentData.ConceptName,
				&currentData.VideoSegmentName,
				&currentData.VideoName,
			)
			check(err)
			allData = append(allData, currentData)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&allData)
	}
}

func setup() {
	db, err = sql.Open("mysql", "root:welcome123@tcp(localhost:3306)/db_bfs_toppr")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/api/v1/syllabus/", getSyllabus)
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func main() {
	setup()
}
