package usecase

import (
	"time"

	"service1/modules/entities/events"
	"service1/modules/entities/models"
)

type producerUser struct {
	eventProducer models.EventProducer
}

func NewProducerServiceUsers(eventProducer models.EventProducer) models.ProducerUser {
	return &producerUser{eventProducer}
}

// UserCreated implements ProducerUser.

func (obj *producerUser) UserCreated(user *models.UserRequest, timeStamp time.Time) error {

	return obj.eventProducer.Produce(events.UserCreatedEvent{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		TimeStamp: timeStamp,
	})
}

// UserUpdated implements ProducerUser.
func (obj *producerUser) UserUpdated(user *models.UserRequest, timeStamp time.Time) error {
	return obj.eventProducer.Produce(events.UserUpdatedEvent{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		TimeStamp: timeStamp,
	})
}

// UserDeleted implements ProducerUser.
func (obj *producerUser) UserDeleted(user uint) error {
	return obj.eventProducer.Produce(events.UserDeletedEvent{
		ID: user,
	})
}
