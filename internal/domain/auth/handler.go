package auth

import (
	"fmt"
	"go-template/internal/domain/user"
	"go-template/pkg/logger"
	"go-template/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}

type AuthHandler struct {
	userRepository *user.UserRepository
}

func NewHandler(repository *user.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepository: repository,
	}
}

func (handler *AuthHandler) Login(ctx *fiber.Ctx) error {
	dto := new(AuthDTO)

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

	user, err := handler.userRepository.FindByEmail(dto.Email)

	if err != nil {
		logger.Error(fmt.Sprintf("User not allowed: %v", err))
		return ctx.Status(401).JSON(fiber.Map{
			"message": "not allowed",
		})
	}

	if !utils.ComparePassword(user.Passsword, dto.Password) {
		logger.Info(fmt.Sprintf("database: %s sent: %s", user.Passsword, dto.Password))
		logger.Info(fmt.Sprintf("Incorrect password tryed: %v", err))
		return ctx.Status(400).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to generate token: %v", err))
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
