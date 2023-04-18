package routes

import (
	"boilerplate/controllers/activityController"
	"boilerplate/controllers/todoController"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "devcode updated",
		})
	})
	activity := api.Group("/activity-groups")
	activity.Get("/", activityController.GetAll)
	activity.Get("/:id", activityController.Get)
	activity.Post("/", activityController.Create)
	activity.Patch("/:id", activityController.Update)
	activity.Delete("/:id", activityController.Delete)

	todo := api.Group("/todo-items")
	todo.Get("/", todoController.GetAll)
	todo.Get("/:id", todoController.Get)
	todo.Post("/", todoController.Create)
	todo.Patch("/:id", todoController.Update)
	todo.Delete("/:id", todoController.Delete)
}
