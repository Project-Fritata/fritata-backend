package apihealth

import "github.com/gofiber/fiber/v3"

type HealthRes struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

type Service string

const (
	Auth  Service = "auth"
	Posts Service = "posts"
	Users Service = "users"
)

func Health(c fiber.Ctx, s Service) error {
	c.Status(fiber.StatusOK)
	return c.JSON(HealthRes{
		Status:  "healthy",
		Service: string(s),
	})
}
