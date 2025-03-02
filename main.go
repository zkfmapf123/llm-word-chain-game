package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/zkfmapf123/go-llm/docs"
	"github.com/zkfmapf123/go-llm/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

var (
	DEFAULT_PORT = "3000"
	APP_NAME     = "example"
	PREFIX       = "api"
	PORT         = os.Getenv("PORT")
)

// @title			go-llm
// @version		1.0
// @description	golang llm server
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	fiber@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "fiber",
		AppName:       APP_NAME,
	})

	app.Use(logger.New())
	app.Get("/ping", handlers.PingHandlers)
	app.Get("/swagger/*", swagger.HandlerDefault)

	// config

	// api
	setRouter(app, PREFIX)

	// static
	// app.Static("/", "./static")

	go func() {
		if err := app.Listen(":" + PORT); err != nil {
			log.Printf("Server error: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("Shutting down server...")
	app.Shutdown()
}

func setRouter(app *fiber.App, prefix string) {
	r := app.Group(prefix)

	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, 끝말잇기 시작!!!")
	})

	// 게임 세션 시작
	r.Get("/session", handlers.SessionHandler)

	// 게임 시작
	r.Post("/play", handlers.UserInputsHandlers)
}
