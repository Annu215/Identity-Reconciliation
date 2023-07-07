package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Route(router fiber.Router) {
	fmt.Println("Hello I'm Working!Routeeee")
	route := router.Group("/user")

	route.Post("/add", Add)
}
