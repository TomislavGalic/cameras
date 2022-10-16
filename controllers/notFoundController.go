package controllers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {

	return c.Render("notFound", fiber.Map{})
}
