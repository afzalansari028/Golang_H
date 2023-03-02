package routes

import (
	"learnapi/handler"

	"github.com/gin-gonic/gin"
)

func MyRouter(r *gin.Engine) {

	//routing
	r.GET("/test", handler.Test)
	r.GET("/courses", handler.GetAllCourses)
	r.GET("/course/:id", handler.GetOneCourse)
	r.POST("/course", handler.CreateOneCourse)
	r.PUT("/course/:id", handler.UpdateOnecourse)
	r.DELETE("/course/:id", handler.DeleteOneCourse)

}
