package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Redirect(ctx *fiber.Ctx) error {
	return ctx.Redirect("/projects")
}
