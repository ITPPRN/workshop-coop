package models

import (
	"time"

	"service1/modules/entities/events"
)

type UserRepository interface {
	CreateUser(string, string) (*User, error)
	UpdateUser(uint, string, string) (*User, error)
	DeleteUser(uint) (*string, error)
}

type UserUsecase interface {
	Register(string, string) (*UserRequest, error)
	UpdateAccount(uint, UserRequest) (*UserRequest, error)
	DeleteAccount(uint) (*string, error)
}

type ProducerUser interface {
	UserCreated(user *UserRequest, time time.Time) error
	UserUpdated(user *UserRequest, timeStamp time.Time) error
	UserDeleted(userid uint) error
}

type EventProducer interface {
	Produce(event events.Event) error
}

type EventHandler interface {
	Handle(toppic string, eventByte []byte)
}

type ConsumerRepository interface {
	CreateUserReadedDog(*UserReadDog) error
}
