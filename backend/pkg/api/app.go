package api

import (
	"github.com/aaronclaydon/whensmybin/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupServer(listen string) {
	webApp := fiber.New()

	// group.Get("version", routes.APIVersion)

	routes.BinsRouter(webApp.Group("/bins"))

	webApp.Listen(listen)
}
