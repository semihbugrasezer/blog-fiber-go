package main

import (
    "log"
    "os"

    "go-blog/database"
    "go-blog/handlers"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Initialize the database
    database.InitDB()

    // Create a new HTML template engine
    engine := html.New("./views", ".html")

    // Create a new Fiber app with the template engine
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    // Serve static files from the "./static" directory
    app.Static("/static", "./static")

    // Define routes and their handlers
    app.Get("/", handlers.GetPosts)
    app.Get("/post/:id", handlers.GetPost)
    app.Post("/post", handlers.CreatePost)

    // Get the port from the environment variable or default to 8000
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000" // Default port
    }

    // Start the Fiber app on the specified port
    log.Fatal(app.Listen(":" + port))
}
