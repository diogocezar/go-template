package user

import (
	"encoding/json"
	"fmt"
	"go-template/internal/infra/queue"
	"go-template/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserDTO struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type UserHandler struct {
	userProducer   *queue.Producer
	userRepository *UserRepository
}

func NewHandler(repository *UserRepository, producer *queue.Producer) *UserHandler {
	return &UserHandler{
		userProducer:   producer,
		userRepository: repository,
	}
}

// CreateUser godoc
//
//	@Summary		Create new user
//	@Description	Create new user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		User	true	"User"
//	@Success		200		{string}	string	"OK"
//	@Failure		400		{string}	error	"Bad Request"
//	@Router			/user													[post]
func (handler *UserHandler) Create(ctx *fiber.Ctx) error {
	dto := new(UserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		logger.Error(fmt.Sprintf("Error parsing body: %v", err))
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		logger.Error(fmt.Sprintf("Error validating body: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := handler.userRepository.Create(dto.Name, dto.Email, dto.Password)

	if err != nil {
		logger.Error(fmt.Sprintf("Error creating user: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		logger.Error(fmt.Sprintf("Error marshalling user: %v", err))
		return ctx.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "error marshalling user"})
	}

	err = queue.Publish("users", string(userJson), handler.userProducer)
	if err != nil {
		logger.Error(fmt.Sprintf("Error publishing user to queue: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": "error trying to send message"})
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(user)
}

// RetriveAllUsers godoc
//
//	@Summary		Retrieve all users
//	@Description	Retrieve all users
//	@Tags			user
//	@Produce		json
//	@Success		200	{array}		User	"List of users"
//	@Failure		422	{string}	error	"Unprocessable Entity"
//	@Router			/user [get]
func (handler *UserHandler) FindAll(ctx *fiber.Ctx) error {
	users, err := handler.userRepository.FindAll()

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting users: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(users)
}

// RetrieveUser godoc
//
//	@Summary		Retrieve a single user
//	@Description	Retrieve a single user
//	@Tags			user
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	User	"User found"
//	@Failure		422	{string}	error	"Unprocessable Entity"
//	@Router			/user/{id} [get]
func (handler *UserHandler) FindOne(ctx *fiber.Ctx) error {
	user, err := handler.userRepository.FindOne(ctx.Params("id"))

	if err != nil {
		logger.Error(fmt.Sprintf("Error getting user: %v with id %s", err, ctx.Params("id")))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(user)
}

// UpdateUser godoc
//
//	@Summary		Update an existing user
//	@Description	Update an existing user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"User ID"
//	@Param			payload	body		UserDTO	true	"User DTO"
//	@Success		200		{object}	User	"User updated"
//	@Failure		400		{string}	error	"Bad Request"
//	@Failure		422		{string}	error	"Unprocessable Entity"
//	@Router			/user/{id} [put]
func (handler *UserHandler) Update(ctx *fiber.Ctx) error {
	dto := new(UserDTO)

	if err := ctx.BodyParser(dto); err != nil {
		logger.Error(fmt.Sprintf("Error parsing body: %v", err))
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "bad request"})
	}

	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		logger.Error(fmt.Sprintf("Error validating body: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	user, err := handler.userRepository.Update(ctx.Params("id"), dto.Name, dto.Email)

	if err != nil {
		logger.Error(fmt.Sprintf("Error updating user: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(user)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user
//	@Tags			user
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{string}	string	"User removed"
//	@Failure		422	{string}	error	"Unprocessable Entity"
//	@Router			/user/{id} [delete]
func (handler *UserHandler) Delete(ctx *fiber.Ctx) error {
	err := handler.userRepository.Delete(ctx.Params("id"))

	if err != nil {
		logger.Error(fmt.Sprintf("Error deleting user: %v", err))
		return ctx.
			Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.
		Status(fiber.StatusOK).
		Send([]byte{})
}
