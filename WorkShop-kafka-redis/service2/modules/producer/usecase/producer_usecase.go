package usecase

import (
	"service2/modules/entities/events"
	"service2/modules/entities/models"
)

type producerUser struct {
	eventProducer models.EventHandlerProduce
}

func NewProducerUsecase(eventProducer models.EventHandlerProduce) models.UseCaseProducer {
	return &producerUser{eventProducer}
}

// UserCreated implements ProducerUser.

func (obj *producerUser) UserReaded(event *events.UserReadedEvent) error {

	return obj.eventProducer.Produce(event)
}
