package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learnapi/db"
	"learnapi/models"
	"log"

	"github.com/gin-gonic/gin"
)

//serve home route
func Test(c *gin.Context) {
	c.JSON(200, "Api testing.....!!")
}

func GetAllCourses(c *gin.Context) {
	fmt.Println("Get all courses")

	db := db.SetupDB()
	rows, err := db.Query("select * from courses")
	if err != nil {
		log.Fatal(err)
	}

	var courses []models.Course

	for rows.Next() {
		var courseId string
		var courseName string
		var price int
		var author string

		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal("Scan error:", err)
		}
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	// fmt.Println(courses)
	c.JSON(200, courses)
}

func GetOneCourse(c *gin.Context) {
	fmt.Println("Get one course")

	// 	// grab id from request
	id := c.Param("id")
	db := db.SetupDB()
	rows, err := db.Query("select * from courses where course_id = $1", id)
	if err != nil {
		log.Fatal("Scan error:", err)
	}
	var course []models.Course
	for rows.Next() {
		var courseId string
		var courseName string
		var price int
		var author string

		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		course = append(course, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	log.Print(course)
	if course != nil {
		c.JSON(200, course)
	} else {
		c.JSON(404, "No data is available of this id")
	}
}

func CreateOneCourse(c *gin.Context) {
	fmt.Println("Create one course")

	var course models.Course
	// unmarshaling the request - converting the json to go object
	jsonDataBytes, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(jsonDataBytes, &course)

	// decoding the requestanother way of converting the json to go object
	// json.NewDecoder(c.Request.Body).Decode(&course)

	//what about - {}
	if course.CourseId == "" && course.CourseName == "" && course.CoursePrice == 0 && course.Author == "" {
		c.JSON(404, "No data inside JSON")
		return
	}

	db := db.SetupDB()
	var courses []models.Course
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
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	for _, dbCourses := range courses { //checking duplicate data
		if courseId == dbCourses.CourseId {
			c.JSON(409, "This course is already available, Please try with other...!")
			return
		}
	}

	_, err := db.Exec("INSERT INTO courses(course_id,course_name,course_price,course_author) values($1,$2,$3,$4)", courseId, courseName, price, author)
	if err != nil {
		log.Print("Error while inserting data", err)
		return
	}
	c.JSON(200, "Course added...")
}

func UpdateOnecourse(c *gin.Context) {
	fmt.Println("Update one course")
	id := c.Param("id")

	var updatecourse models.Course
	json.NewDecoder(c.Request.Body).Decode(&updatecourse)
	var (
		courseName   = updatecourse.CourseName
		coursePrice  = updatecourse.CoursePrice
		courseAuthor = updatecourse.Author
	)
	// fmt.Println(courseName, coursePrice, courseAuthor)
	// fmt.Println(id)
	db := db.SetupDB()
	_, err := db.Exec("UPDATE courses SET course_name=$1,course_price=$2,course_author=$3 where course_id=$4", courseName, coursePrice, courseAuthor, id)
	if err != nil {
		log.Print("Error while updating data", err)
		return
	}
	c.JSON(200, "Course updated")
}

func DeleteOneCourse(c *gin.Context) {
	fmt.Println("Delete one course from db")
	id := c.Param("id")
	db := db.SetupDB()

	rows, _ := db.Query("SELECT * FROM courses")
	var courses []models.Course
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
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}

	for _, courseDB := range courses {
		if courseDB.CourseId == id {
			_, err := db.Exec("DELETE FROM courses where course_id = $1 ", id)
			if err != nil {
				log.Print("Error while deleting data", err)
				return
			}
			c.JSON(200, "Course is deleted of given id")
			return
		}
	}
	c.JSON(200, "Course not found by given id")
}
