package main

import (
	"database/sql"
	"final-project/controllers"
	"final-project/database"
	"final-project/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "final-project"
// )

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed: ", err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})
	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)

	router.GET("/task-status", controllers.GetAllTaskStatus)
	router.POST("/task-status", controllers.InsertTaskStatus)
	router.PUT("/task-status/:id", controllers.UpdateTaskStatus)
	router.DELETE("/task-status/:id", controllers.DeleteTaskStatus)

	router.GET("/task", controllers.GetAllTask)
	router.POST("/task", middleware.MiddlewareAuth(), controllers.InsertTask)
	router.PUT("/task/:id", middleware.MiddlewareAuth(), controllers.UpdateTask)
	router.DELETE("/task/:id", controllers.DeleteTask)

	router.GET("/project", controllers.GetAllProject)
	router.POST("/project", middleware.MiddlewareAuth(), controllers.InsertProject)
	router.PUT("/project/:id", middleware.MiddlewareAuth(), controllers.UpdateProject)
	router.DELETE("/project/:id", controllers.DeleteProject)

	// router.Run("localhost:8080")
	router.Run(":" + os.Getenv("PORT"))
}
