package main

import (
	"fmt"
	"github.com/davidyapcc/todo-app-backend/handlers"
	"github.com/davidyapcc/todo-app-backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	db := InitDB()
	if db == nil {
		panic("Failed to initialize database")
	}
	defer CloseDB()

	handlers.SetDB(db)
	routes.InitTodoRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server", err)
		panic(err)
	}
}
