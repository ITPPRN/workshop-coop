package events

import (
	"encoding/json"
	"time"
)

type Event interface {
	String() string
}

var SubscribedTopics = []string{
	UserCreatedEvent{}.String(),
	UserUpdatedEvent{}.String(),
	UserDeletedEvent{}.String(),
}

type UserCreatedEvent struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	TimeStamp time.Time `json:"timeStamp"`
}

type UserUpdatedEvent struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	TimeStamp time.Time `json:"timeStamp"`
}

type UserDeletedEvent struct {
	ID uint `json:"id"`
}

type UserReadedEvent struct {
	UserId     uint            `json:"userId"`
	DogId      uint            `json:"dogId"`
	DogDetails json.RawMessage `json:"dogDetails"`
	TimeStamp  time.Time       `json:"timeStamp"`
}

func (UserCreatedEvent) String() string {
	return "UserCreated"
}
func (UserUpdatedEvent) String() string {
	return "UserUpdated"
}
func (UserReadedEvent) String() string {
	return "UserReaded"
}
func (UserDeletedEvent) String() string {
	return "UserDeleted"
}
