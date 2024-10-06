package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type Repository struct {
	DB *gorm.DB
}

func (r * Repository) SetupRoutes(app *fiber.Application){
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBooks)
	api.Delete("/delete_books/:id", r.DeleteBooks)
	api.Get("/get_books/:id", r.GetBooks)
} 

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	db, err := storage.NewConnection(config)

	if err != nil{
		log.Fatal("could not load the database")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
