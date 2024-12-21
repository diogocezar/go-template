package user

import (
	"encoding/json"
	"go-template/internal/infra/queue"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"required,email"`
}

type FilterUserDTO struct {
	ID string `json:"id" validate:"required"`
}

type Controller struct {
	producer   *queue.Producer
	repository *Repository
}

func MakeController(repository *Repository, producer *queue.Producer) *Controller {
	return &Controller{
		repository: repository,
		producer:   producer,
	}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	dto := new(UserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		log.Println(err)
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
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

	userJSON, err := json.Marshal(user)
	if err != nil {
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "error marshalling user"})
	}

	err = queue.Publish("users", string(userJSON), c.producer)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": "error trying to send message"})
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(user)
}

func (c *Controller) FindAll(ctx *fiber.Ctx) error {

	users, err := c.repository.FindAll()

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(users)
}

func (c *Controller) FindOne(ctx *fiber.Ctx) error {
	dto := FilterUserDTO{
		ID: ctx.Params("id"),
	}

	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.repository.FindOne(dto.ID)

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(user)
}

func (c *Controller) Update(ctx *fiber.Ctx) error {
	dto := new(UserDTO)
	filter := FilterUserDTO{
		ID: ctx.Params("id"),
	}

	if err := ctx.BodyParser(dto); err != nil {
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := c.repository.Update(filter.ID, dto.Name, dto.Email)

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(user)
}

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	dto := FilterUserDTO{
		ID: ctx.Params("id"),
	}

	err := c.repository.Delete(dto.ID)

	if err != nil {
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		Send([]byte{})
}
