package controllers

import (
	"auth-system/app/core/constants"
	"auth-system/app/domain"
	"auth-system/app/usecases"
	"github.com/gofiber/fiber/v2"
	"math"
	"strconv"
)

type DividaController struct {
	usecase usecases.DividaUsecase
}

func NewDividaController(u usecases.DividaUsecase) *DividaController {
	return &DividaController{
		usecase: u,
	}
}

func (dc *DividaController) CriarDivida(c *fiber.Ctx) error {

	empresaIDStr := c.Locals(constants.UserSessionKey)
	empresaID, err := strconv.ParseUint(empresaIDStr.(string), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "usuário não autenticado ou inválido",
		})
	}

	var payload struct {
		CPFConsumidor string        `json:"cpfConsumidor"`
		Divida        domain.Divida `json:"divida"`
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "dados inválidos",
		})
	}

	createdDivida, err := dc.usecase.CriarDivida(payload.Divida, payload.CPFConsumidor, uint(empresaID))
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdDivida)
}

func (dc *DividaController) ListarDividas(c *fiber.Ctx) error {
	userIDStr := c.Locals(constants.UserSessionKey) // ajusta conforme como está salvando na sessão
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "usuário não autenticado ou inválido",
		})
	}
	// Chama o usecase para listar as dívidas
	dividas, err := dc.usecase.ListarDividas(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dividas)
}

func (dc *DividaController) CalcularScore(c *fiber.Ctx) error {
	userIDStr := c.Locals(constants.UserSessionKey)
	userID, err := strconv.ParseUint(userIDStr.(string), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "usuário não autenticado ou inválido",
		})
	}

	score, err := dc.usecase.AtualizarEObterScore(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	truncatedScore := math.Floor(score*100) / 100
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"score": truncatedScore,
	})
}
