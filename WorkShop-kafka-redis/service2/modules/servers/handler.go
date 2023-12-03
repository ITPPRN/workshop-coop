package servers

import (
	"github.com/gofiber/fiber/v2"

	consumerHandler "service2/modules/consumer/handlers"
	consumerUsecase "service2/modules/consumer/usecase"
	"service2/modules/dog/controller"
	"service2/modules/dog/repository"
	"service2/modules/dog/usecase"
	userRepo "service2/modules/users/repository"
	_handlerProducer "service2/modules/producer/handlers"
	_publisherUsecase "service2/modules/producer/usecase"
)

func (s *server) Handlers() error {

	//repo dog
	dogRepository := repository.NewDogRepositoryDB(s.Db)
	userRepo := userRepo.NewUserRepositoryDB(s.Db)

	// consumer
	consumeUsecase := consumerUsecase.NewConsumerUsecase(userRepo)
	eventHandlerConsumer := consumerHandler.NewEventHandler(consumeUsecase)
	s.consumerGroupHandler = consumerHandler.NewHandlerConsumeGroup(eventHandlerConsumer)

	// producer
	handlerProducer := _handlerProducer.NewEventHandlerProducer(s.SyncProducer)
	producerUsecase := _publisherUsecase.NewProducerUsecase(handlerProducer)

	v1 := s.App.Group("/v1")
	dogUsecase := usecase.NewDogService(dogRepository,userRepo, s.Redis, producerUsecase)
	controller.NewDogController(v1, dogUsecase)

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
