package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"service1/modules/entities/models"
)

type userHandler struct {
	userSrv models.UserUsecase
}

func NewUserController(router fiber.Router, userSrv models.UserUsecase) {
	controllers := &userHandler{
		userSrv: userSrv,
	}
	router.Post("/register", controllers.Register)
	router.Put("/update/:id", controllers.Update)
	router.Delete("/delete/:id", controllers.Delete)
}

func (h userHandler) Register(c *fiber.Ctx) error {
	// รับค่าตัวแปรจาก body ของ HTTP POST
	request := models.UserRequest{}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	response, err := h.userSrv.Register(request.Name, request.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": response,
	})
}

func (h userHandler) Update(c *fiber.Ctx) error {

	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	// รับค่าตัวแปรจาก body ของ HTTP PUT
	request := models.UserRequest{}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	response, err := h.userSrv.UpdateAccount(uint(userID), request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": response,
	})
}

func (h userHandler) Delete(c *fiber.Ctx) error {

	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	response, err := h.userSrv.DeleteAccount(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": response,
	})
}
