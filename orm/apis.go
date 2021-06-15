package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"gorm.io/gorm"
)

var Syl []Syllabus

func ord(db *gorm.DB) *gorm.DB {
	return db.Order("id ASC").Where("is_active = ?", 1)
}

func PreloadSyllabus(db *gorm.DB) {
	db.Preload("Subjects", ord).Preload("Subjects.Topics", ord).Preload("Subjects.Topics.Subtopics", ord).Find(&Syl, "is_active=?", 1)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Bisearch(arr ...Searcher) {
	fmt.Println(arr[0].GetID())
}

func GetSyllabus(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		req.ParseForm()

		id, err := strconv.Atoi(req.FormValue("id"))
		check(err)
		fmt.Println(id)
		var allData []*Syllabus
		for _, u := range Syl {
			allData = append(allData, &Syllabus{
				Id:   u.Id,
				Name: u.Name,
			})
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&allData)
	}
}

func GetSubjects(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		req.ParseForm()

		id, err := strconv.Atoi(req.FormValue("id"))
		check(err)
		var allData []*Subject

		idx := sort.Search(len(Syl), func(i int) bool {
			return Syl[i].Id >= id
		})
		if idx >= len(Syl) {
			http.Error(w, "Syllabus Doesn't Exist", http.StatusNotFound)
			return
		}
		if Syl[idx].Id == id {
			for _, t := range Syl[idx].Subjects {
				allData = append(allData, &Subject{
					Id:    t.Id,
					Name:  t.Name,
					SylId: Syl[idx].Id,
				})
			}
		} else {
			http.Error(w, "Syllabus Doesn't Exist", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&allData)
	}
}
