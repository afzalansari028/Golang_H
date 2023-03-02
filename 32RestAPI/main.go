package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	DB_USER     = "postgres"
// 	DB_PASSWORD = "king"
// 	DB_NAME     = "library"
// )

// // DB set up
// func setupDB() *sql.DB {
// 	db, err := sql.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

// //Model for course - file
// type Course struct {
// 	CourseId   string `json:"courseid"`
// 	CourseName string `json:"coursename"`
// 	// CoursePrice int     `json:"-"`
// 	CoursePrice int    `json:"price"`
// 	Author      string `json:"author"`
// }

//middleware, helper - file
// func (c *Course) IsEmpty() bool {
// 	return c.CourseId == "" && c.CourseName == ""
// }

func main() {
	fmt.Println("API - Learning in Golang")
	router := mux.NewRouter()
	// fmt.Println(setupDB())

	//routing
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOnecourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":5000", router))

}

//serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the API by golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	db := setupDB()
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	id := params["id"]

	db := setupDB()
	rows, err := db.Query("select * from courses where course_id = ?", id)
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
	// log.Print(courses)
	if courses != nil {
		json.NewEncoder(w).Encode(courses)
	} else {
		json.NewEncoder(w).Encode("No data is available of this id")
	}
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what about - {}
	var course Course
	var courses []Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	db := setupDB()
	var courseId = course.CourseId
	var courseName = course.CourseName
	var price = course.CoursePrice
	var author = course.Author
	// fmt.Println(courseId, courseName, price, author)

	//validate duplicate data
	rows, _ := db.Query("SELECT * FROM courses")
	for rows.Next() {
		var (
			courseId   string
			courseName string
			price      int
			author     string
		)
		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	for _, dbCourses := range courses {
		if courseId == dbCourses.CourseId {
			json.NewEncoder(w).Encode("This course is already available, Please try with other...!")
			return
		}
	}

	db.Query("INSERT INTO courses(course_id,course_name,course_price,author) values(?,?,?,?)", courseId, courseName, price, author)
	json.NewEncoder(w).Encode("Course added...")
	w.WriteHeader(http.StatusOK)
}

func updateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	// w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var updatecourse Course
	json.NewDecoder(r.Body).Decode(&updatecourse)
	var (
		courseName   = updatecourse.CourseName
		coursePrice  = updatecourse.CoursePrice
		courseAuthor = updatecourse.Author
	)
	fmt.Println(courseName, coursePrice, courseAuthor)
	db := setupDB()
	db.Query("UPDATE courses SET course_name=?,course_price=?,author=? where course_id=?", courseName, coursePrice, courseAuthor, id)
	json.NewEncoder(w).Encode("Course updated")
	w.WriteHeader(http.StatusOK)
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course from db")
	w.Header().Set("Content-Type,", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	db := setupDB()

	rows, _ := db.Query("SELECT * FROM courses")
	var courses []Course
	for rows.Next() {
		var (
			courseId   string
			courseName string
			price      int
			author     string
		)
		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}

	for _, courseDB := range courses {
		if courseDB.CourseId == id {
			db.Query("DELETE FROM courses where course_id = ? ", id)
			json.NewEncoder(w).Encode("Course is deleted of given id")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found by given id")
	w.WriteHeader(http.StatusNotFound)
}
