package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"service2/modules/entities/models"
)

type dogHandler struct {
	dogUsecase models.DogUsecase
}

func NewDogController(router fiber.Router, dogUsecase models.DogUsecase) {
	controllers := &dogHandler{
		dogUsecase: dogUsecase,
	}

	router.Get("/dogs", controllers.getAllDogData)
	router.Get("/dog/:userId/:dogId", controllers.UserReaded)

}
func (h *dogHandler) getAllDogData(c *fiber.Ctx) error {
	m, err := h.dogUsecase.GetDogs()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			models.ResponseError{
				Message:    err.Error(),
				Status:     fiber.ErrNotFound.Message,
				StatusCode: fiber.ErrNotFound.Code,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		models.ResponseData{
			Message:    "Succeed",
			Status:     "OK",
			StatusCode: fiber.StatusOK,
			Data:       m,
		},
	)
}

func (h *dogHandler) UserReaded(c *fiber.Ctx) error {
	dogIdParam := c.Params("dogId")
	userIdParam := c.Params("userId")
	if userIdParam == "" || dogIdParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.ResponseError{
				Message:    "Require parameters",
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.ErrBadRequest.Code,
			},
		)
	}
	dogid, err := strconv.ParseUint(dogIdParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.ResponseError{
				Message:    "Invalid parameter type",
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.ErrBadRequest.Code,
			},
		)
	}
	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			models.ResponseError{
				Message:    "Invalid parameter type",
				Status:     fiber.ErrBadRequest.Message,
				StatusCode: fiber.ErrBadRequest.Code,
			},
		)
	}
	res, err := h.dogUsecase.UserReadData(uint(userId), uint(dogid))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(
			models.ResponseError{
				Message:    err.Error(),
				Status:     fiber.ErrNotFound.Message,
				StatusCode: fiber.ErrNotFound.Code,
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		models.ResponseData{
			Message:    "Succeed",
			Status:     "OK",
			StatusCode: fiber.StatusOK,
			Data:       res,
		},
	)
}
