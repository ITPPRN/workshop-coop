package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"

	"service1/modules/consumer/usecase"
	"service1/modules/entities/events"
	"service1/modules/entities/models"
)

type eventHandler struct {
	consumer usecase.ConsumerService
}

func NewEventHandler(consumer usecase.ConsumerService) models.EventHandler {
	return &eventHandler{consumer}
}

func (obj *eventHandler) Handle(topic string, eventBytes []byte) {
	log.Info("consume topic:", topic)
	switch topic {
	case events.UserReadedEvent{}.String():
		event := events.UserReadedEvent{}
		err := json.Unmarshal(eventBytes, &event)
		if err != nil {
			log.Error(err)
			return
		}
		log.Debug("topic: ", topic, "user id:", event.UserId, "dog id:", event.DogId)
		err = obj.consumer.UserReaded(event)
		if err != nil {
			log.Error(err)
			return
		}
	}
	return

}
