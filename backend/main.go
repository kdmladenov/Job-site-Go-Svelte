package main

import (
	"jobs-site/controllers"
	"jobs-site/initializers"
	"jobs-site/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", middleware.RequireAuthentication, controllers.Logout)

	r.GET("/jobs", controllers.GetJobs)
	r.GET("/jobs/:id", controllers.GetJob)
	r.POST("/jobs", middleware.RequireAuthentication, middleware.ValidateJobBody, controllers.CreateJob)
	r.PUT("/jobs/:id", middleware.RequireAuthentication, middleware.RequireAuthorizationJob, middleware.ValidateJobBody, controllers.UpdateJob) 
	r.DELETE("/jobs/:id", middleware.RequireAuthentication, middleware.RequireAuthorizationJob, controllers.DeleteJob)

	r.Run()
}
