package main

import (
	"embed"
	"fmt"
	"net/http"
	"publisher/config"
	"publisher/services"
	"publisher/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func main() {
	c := config.Read()
	startServer(
		&services.Context{
			Config: c,
			DB:     utils.ConnectSqlite(fmt.Sprintf("%s.db", c.DatabaseName), &gorm.Config{}),
		},
	)
}

func startServer(sc *services.Context) {
	app := fiber.New()
	enabledLog(app)
	embedAssets(app)
	registerApi(app, sc)
	app.Listen(utils.HostString(sc.Config.Host, sc.Config.Port))
}

func enabledLog(app *fiber.App) {
	app.Use(logger.New())
}

//go:embed static/*
var embedDirStatic embed.FS

func embedAssets(app *fiber.App) {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		Browse:     true,
		PathPrefix: "/static",
		Index:      "index.html",
	}))
}

func registerApi(app *fiber.App, sc *services.Context) {
	api := app.Group("/api")
	api.Get("/publish", services.Publish(sc))
}
