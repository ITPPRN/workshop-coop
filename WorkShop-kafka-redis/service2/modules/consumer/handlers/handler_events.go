package handlers

import (
	"encoding/json"
	"service2/modules/entities/events"
	"service2/modules/entities/models"

	"github.com/gofiber/fiber/v2/log"
)

type eventHandler struct {
	consumer models.ConsumerUsecase
}

func NewEventHandler(consumer models.ConsumerUsecase) models.EventHandlerConsume {
	return &eventHandler{consumer}
}

func (obj *eventHandler) Handle(topic string, eventBytes []byte) {
	log.Info("consume topic:", topic)
	switch topic {
	case events.UserCreatedEvent{}.String():
		event := events.UserCreatedEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Error(err)
			return
		}
		err = obj.consumer.UserCreated(event)
		if err != nil {
			log.Error(err)
			return
		}
	case events.UserUpdatedEvent{}.String():
		event := events.UserUpdatedEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Error(err)
			return
		}
		err = obj.consumer.UserUpdated(event)
		if err != nil {
			log.Error(err)
			return
		}
	case events.UserDeletedEvent{}.String():
		event := events.UserDeletedEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Error(err)
			return
		}
		err = obj.consumer.UserDeleted(event)
		if err != nil {
			log.Error(err)
			return
		}

	}
}
