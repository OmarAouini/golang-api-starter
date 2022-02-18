package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

var serverAddress = "0.0.0.0:8080"
var DB gorm.DB
var router *fiber.App

func init() {

	//logger config
	ConfigLogger()
	// env config
	ConfigEnv()

	//fiber router
	router = fiber.New()
	router.Use(logger.New())
	router.Use(recover.New())
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, Api-Secret",
		AllowMethods: "GET, POST, PUT, DELETE, HEAD, OPTIONS",
	}))

	//routes
	v2_1 := router.Group("/api/v2.1") // this is for versioning, use for avoiding api version clash
	v2_1.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("OK")
	})

	//normal endpoints
	restreamers := v2_1.Group("/restreamers")
	restreamers.Get("/", JwtProtected(), GetRestreamers)
	restreamers.Post("/", CreateRestreamer)
	restreamers.Put("/", UpdateRestreamer)
	events := restreamers.Group("/events")
	events.Get("/", EventsHistory)
	events.Post("/", CreateEvent)
	events.Put("/:event_id/start", StarEvent)
	events.Put("/:event_id/finish", SetEventCompleted)

	// customers := v2_1.Group("/customers")

	videos := v2_1.Group("/videos")
	videos.Post("/", JwtProtected(), GetVideos)

	//secured endpoints
	secured := v2_1.Group("/secured")
	restreamers_secured := secured.Group("/restreamers")
	restreamers_secured.Post("/", GetRestreamersWithApiKey)
	// customers_secured := secured.Group("/customers")
	videos_secured := secured.Group("/videos")
	videos_secured.Post("/", GetVideosWithApiKey)

}

func main() {
	fmt.Println()
	fmt.Printf("\nENV: %s", CONFIG.AppEnv)
	fmt.Printf("\nRUNNING MODE: %s", CONFIG.RunningMode)
	router.Listen(serverAddress)
}
