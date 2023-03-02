package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// func (c *Course) IsEmpty() bool {
// 	return c.CourseId == "" && c.CourseName == ""
// }

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	db := ConnectionProvider.GetConnection()
	rows, err := db.Query("select * from courses")
	if err != nil {
		log.Fatal(err)
	}

	var courses []Course

	for rows.Next() {
		var courseId string
		var courseName string
		var price int
		var author string

		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	// fmt.Println(courses)
	json.NewEncoder(w).Encode(courses)
}
