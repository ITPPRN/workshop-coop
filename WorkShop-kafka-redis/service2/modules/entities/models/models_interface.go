package models

import (
	"encoding/json"

	"service2/modules/entities/events"
)

type DogRepository interface {
	CreateDog(*Dog) error
	GetDogs() ([]Dog, error)
	HasDog() (bool, error)
	DogExists(dogID uint) bool
	FindDogByID(dogID uint) (*Dog, error)
}

type DogUsecase interface {
	GetDogs() ([]Dog, error)
	UserReadData(uint,uint) (json.RawMessage, error)
}

type EventHandlerProduce interface {
	Produce(event events.Event) error
}

type EventHandlerConsume interface {
	Handle(toppic string, eventByte []byte)
}

type UseCaseProducer interface {
	UserReaded(user *events.UserReadedEvent) error
}

type ConsumerUsecase interface {
	UserCreated(event events.UserCreatedEvent) error
	UserUpdated(event events.UserUpdatedEvent) error
	UserDeleted(event events.UserDeletedEvent) error
}

type UserRepository interface {
	CreateUser(string, string) (*User, error)
	UpdateUser(uint, string, string) (*User, error)
	DeleteUser(uint) (*string, error)
	UserExists(userID uint) bool
}
