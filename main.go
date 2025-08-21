package main

import (
	"log"
	"gotasker/database"
	"gotasker/handlers"
	"gotasker/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "GoTasker API is running"})
	})

	public := router.Group("/api")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/tasks", handlers.GetTasks)
		protected.GET("/tasks/:id", handlers.GetTask)
		protected.POST("/tasks", handlers.CreateTask)
		protected.PUT("/tasks/:id", handlers.UpdateTask)
		protected.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	log.Println("🚀 GoTasker API starting on :8080")
	log.Println("📝 Available endpoints:")
	log.Println("   POST /api/register - Register a new user")
	log.Println("   POST /api/login - Login user")
	log.Println("   GET /api/tasks - Get all tasks (auth required)")
	log.Println("   GET /api/tasks/:id - Get specific task (auth required)")
	log.Println("   POST /api/tasks - Create new task (auth required)")
	log.Println("   PUT /api/tasks/:id - Update task (auth required)")
	log.Println("   DELETE /api/tasks/:id - Delete task (auth required)")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
