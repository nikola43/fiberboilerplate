package controllers

import (
	"github.com/gofiber/fiber/v2"
	requestmodels "github.com/nikola43/fiberboilerplate/internal/api/v1/models/requests"
	"github.com/nikola43/fiberboilerplate/internal/api/v1/services"
	"github.com/nikola43/fiberboilerplate/pkg/utils"
)

func Login(context *fiber.Ctx) error {
	request := new(requestmodels.LoginUserRequest)

	err := utils.ParseAndValidate(context, request)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	loginResponse, err := services.LoginClient(request.Email, request.Password)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusNotFound, err)
	}

	return context.JSON(loginResponse)
}

func Signup(context *fiber.Ctx) error {
	request := new(requestmodels.SignupUserRequest)

	err := utils.ParseAndValidate(context, request)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	err = services.SignUpClient(request)
	if err != nil {
		return utils.ReturnErrorResponse(context, fiber.StatusBadRequest, err)
	}

	return utils.ReturnSuccessResponse(context)
}
