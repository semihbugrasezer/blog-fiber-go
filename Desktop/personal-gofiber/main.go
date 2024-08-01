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
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    database.InitDB()

    engine := html.New("./views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Static("/static", "./static")

    app.Get("/", handlers.GetPosts)
    app.Get("/post/:id", handlers.GetPost)
    app.Post("/post", handlers.CreatePost)

    log.Fatal(app.Listen(os.Getenv("PORT")))
}
