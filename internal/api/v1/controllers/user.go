package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/services"
	"github.com/nikola43/fiberboilerplate/pkg/utils"
)

// GetUserById returns a user by id
func GetUserById(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	user, err := services.GetUserById(id)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusNotFound, err)
	}

	return context.JSON(user)
}

// Update updates a user by id
func Update(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	user, err := services.GetUserById(id)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusNotFound, err)
	}

	if err := context.BodyParser(user); err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	user, err = services.Update(user)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusInternalServerError, err)
	}

	return context.JSON(user)
}

// Delete deletes a user by id
func Delete(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 10, 64)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	err = services.DeleteUser(id)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusInternalServerError, err)
	}

	return context.SendStatus(fiber.StatusNoContent)
}
