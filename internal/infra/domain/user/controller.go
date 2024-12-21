package user

import (
	"errors"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Controller struct {
	repository *Repository
}

func MakeController(repository *Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	dto := new(UserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	// TODO: move this validation to a separate structure
	if err := func() error {
		if dto.Name == "" {
			return errors.New("missing name")
		}

		if len(dto.Name) > 100 {
			return errors.New("invalid name")
		}

		if dto.Email == "" {
			return errors.New("missing email")
		}

		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		re := regexp.MustCompile(emailRegex)
		if !re.MatchString(dto.Email) {
			return errors.New("invalid email")
		}
		return nil
	}(); err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.repository.Create(dto.Name, dto.Email)

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(user)
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (c *Controller) FindOne(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	panic("unimplemented")
}
