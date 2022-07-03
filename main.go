package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"main/controller"
	"main/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	router := gin.Default()
	router.Use(middleware.HandleCrossOriginRequest())

	connStr := os.Getenv("DATABASE_URL")

	if len(connStr) == 0 {
		connStr = "postgres://badrri:secret123@localhost:5432/test_db?sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error", err)
		panic("Error connecting to DB ")
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("The server is up and running at port %v", port))
	})

	router.GET("/people/all", func(ctx *gin.Context) { controller.GetListOfPeople(ctx, db) })
	router.POST("/people/add", func(ctx *gin.Context) { controller.AddPeople(ctx, db) })

	router.Run(fmt.Sprintf(":%v", port))
}
