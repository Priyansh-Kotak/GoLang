package main

import (
	"Project/models"
	"Project/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBooks(context *fiber.Ctx) error {
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "reqest failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"messgae": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book created successfully"})

	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "books retrieved successfully", "data": bookModels})

	return nil
}


func (r *Repository) DeleteBooks(context *fiber.Ctx) error {
	bookModels := models.Book{}
	id := context.Params("id")

	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message" : "id does not found",
		})		
		return nil
	}

	err := r.DB.Delete(bookModels, id)
	
	if err!= nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message" : "Failed to delete the book",
		})

		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "Book deleted successfully",
	})

	return nil

}

func (r *Repository) GetBooksById(context *fiber.Ctx) error{
	id := context.Params("id")

	bookModels := models.Book{}

	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message" :"Book id is empty",
		})
		return nil
	}

	fmt.Println("Id we received: ", id)

	err := r.DB.Where("id = ?", id).First(bookModels).Error

	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message" : "Book not found",
		})
		return err 
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "Book retrieved successfully",
        "data" : bookModels,
	})

	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBooks)
	api.Delete("/delete_books/:id", r.DeleteBooks)
	api.Get("/books", r.GetBooks)
	api.Get("/get_books/:id", r.GetBooksById)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateBook(db)
	if err != nil {
		log.Fatal("could not migrate the db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
