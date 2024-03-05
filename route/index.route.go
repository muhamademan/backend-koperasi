package route

import (
	"backend-koperasi/config"
	"backend-koperasi/handler"
	"backend-koperasi/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	// r.Get("/user", handler.UserHandlerRead)

	r.Post("/login", handler.LoginHandler)

	r.Static("/public", config.ProjectRootPath+"/public/assets")
	r.Get("/users", middleware.Auth, handler.GetAllUsers)
	r.Post("/users", middleware.Auth, handler.CreateUser)
	r.Get("/users/:id", middleware.Auth, handler.GetById)
	r.Put("/users/:id", middleware.Auth, handler.UpdateUser)
	r.Delete("/users/:id", middleware.Auth, handler.DeleteUser)
}
