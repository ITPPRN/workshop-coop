package servers

import (
	"github.com/gofiber/fiber/v2"

	_eventHandler "service1/modules/consumer/handlers"
	consumerRepo "service1/modules/consumer/repositories"
	consumerUsecase "service1/modules/consumer/usecase"
	"service1/modules/producer/handlers"
	_producerUsecase "service1/modules/producer/usecase"
	"service1/modules/users/controller"
	"service1/modules/users/repository"
	"service1/modules/users/usecase"
)

func (s *server) Handlers() error {

	// Kafka
	cosumeRepo := consumerRepo.NewsConsumerRepository(s.Db)
	consumeRepo := consumerUsecase.NewConsumerUsecase(cosumeRepo)
	eventHandler := _eventHandler.NewEventHandler(consumeRepo)
	s.consumerGroupHandler = _eventHandler.NewHandlerConsumeGroup(eventHandler)

	// Group a version
	v1 := s.App.Group("/v1")

	//user
	usersGroup := v1.Group("/user")
	userRepository := repository.NewUserRepositoryDB(s.Db)
	publisherHandler := handlers.NewEventProducer(s.SyncProducer)
	producerUsecase := _producerUsecase.NewProducerServiceUsers(publisherHandler)
	dogUsecase := usecase.NewUserService(userRepository, producerUsecase)
	controller.NewUserController(usersGroup, dogUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil

}
