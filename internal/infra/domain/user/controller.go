package user

import "github.com/gofiber/fiber"

type Controller struct {
	repository *Repository
}

func MakeController(repository *Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) FindOne(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	return nil
}
