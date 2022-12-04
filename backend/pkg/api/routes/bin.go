package routes

import "github.com/gofiber/fiber/v2"

func BinsRouter(router fiber.Router) {
	router.Get("/", listJourneys)
	router.Get("/:district/:namenumber/:postcode", getBin)
}

func listJourneys(c *fiber.Ctx) error {
	c.SendString("NOT IMPLEMENTED")
	return nil
}

func getBin(c *fiber.Ctx) error {
	district := c.Params("district")
	nameNumber := c.Params("nameNumber")
	postCode := c.Params("postcode")

	c.SendStatus(503)
	return c.JSON(fiber.Map{
		"district":   district,
		"nameNumber": nameNumber,
		"postCode":   postCode,
	})
}
