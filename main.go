package main

import (
	"fiber-mongo/app"
	"fiber-mongo/db"
	ihttp "fiber-mongo/http"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	args := os.Args
	op := "server"
	if len(args) > 1 {
		op = args[0]
	}
	if err := run(op); err != nil {
		fmt.Println(fmt.Errorf("error - server failed to start. err: %v", err))
	}
}

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func run(op string) error {
	d, err := db.NewMongoStore()
	if err != nil {
		return errors.Wrap(err, "unable to connect to db!!")
	}
	svc := app.NewUserSvc(d)

	application := fiber.New(fiber.Config{
		//JSONEncoder: json.Marshal,
		//JSONDecoder: json.Unmarshal,
	})
	application.Use(logger.New())
	h := ihttp.NewHandler(svc)
	apiGroup := application.Group("api")
	ihttp.Routes(apiGroup, h)
	port := Getenv("PORT", "8080")
	log.Println("port", port)
	host := Getenv("HOST", "0.0.0.0")
	result := fmt.Sprintf("%s:%s", host, port)
	return application.Listen(result)
}
