package controllers

import (
	"net/http"

	"cashier_indrajaya/models"
	"cashier_indrajaya/services"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service}
}

func (ctrl *TransactionController) Create(c *fiber.Ctx) error {
	var tx models.Transaction
	if err := c.BodyParser(&tx); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := ctrl.service.Create(&tx); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(tx)
}

func (ctrl *TransactionController) GetAll(c *fiber.Ctx) error {
	transactions, err := ctrl.service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(transactions)
}
