package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/RohitKuwar/go_fiber2/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	
	app.Use(logger.New())
	routes.Setup(app)
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  port := os.Getenv("PORT")

	fmt.Println("Server is runnig on port:", port)
	app.Listen(":"+port)
	
}