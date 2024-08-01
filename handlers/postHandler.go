package handlers

import (
	"database/sql"
	"go-blog/database"
	"go-blog/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
    rows, err := database.DB.Query("SELECT id, title, content FROM posts")
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    defer rows.Close()

    posts := []models.Post{}
    for rows.Next() {
        var post models.Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
            return c.Status(500).SendString(err.Error())
        }
        posts = append(posts, post)
    }

    return c.Render("index", fiber.Map{
        "Posts": posts,
    })
}

func GetPost(c *fiber.Ctx) error {
    id := c.Params("id")
    postID, err := strconv.Atoi(id)
    if err != nil {
        return c.Status(400).SendString("Invalid post ID")
    }

    var post models.Post
    row := database.DB.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", postID)
    if err := row.Scan(&post.ID, &post.Title, &post.Content); err != nil {
        if err == sql.ErrNoRows {
            return c.Status(404).SendString("Post not found")
        }
        return c.Status(500).SendString(err.Error())
    }

    return c.Render("post", fiber.Map{
        "ID":      post.ID,
        "Title":   post.Title,
        "Content": post.Content,
    })
}

func CreatePost(c *fiber.Ctx) error {
    type Request struct {
        Title   string `json:"title"`
        Content string `json:"content"`
    }

    var req Request
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).SendString(err.Error())
    }

    query := "INSERT INTO posts (title, content) VALUES (?, ?)"
    result, err := database.DB.Exec(query, req.Title, req.Content)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    id, err := result.LastInsertId()
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }

    return c.JSON(fiber.Map{
        "id": id,
    })
}
