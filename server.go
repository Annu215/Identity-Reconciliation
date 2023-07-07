package main

import (
	"TASK/router"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

func main() {
	connect()
	app := fiber.New(
		fiber.Config{
			// ErrorHandler:  ErrorHandler,
			// JSONEncoder:   object.Encode,
			StrictRouting: true,
			CaseSensitive: true,
		},
	)
	router.Configure(app)
	fmt.Println("Hello I'm Working!")

	err := app.Listen(":4200")
	log.Fatalln(err)
}

//	func ErrorHandler(c *fiber.Ctx, err error) error {
//		code := fiber.StatusInternalServerError
//		if e, ok := err.(*fiber.Error); ok {
//			code = e.Code
//		}
//		response := &Response{
//			Ok:    false,
//			Error: err.Error(),
//		}
//		return c.Status(code).JSON(response)
//	}
func connect() {
	// Set up the connection string
	connStr := "host=localhost port=5432 user=postgres password=CCnew_password dbname=SampleDB sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	// Perform database operations...
}
