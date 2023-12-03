package usecase

import (
	"errors"

	"service2/logs"
	"service2/modules/entities/events"
	"service2/modules/entities/models"
)

type consumerUsecase struct {
	userRepo models.UserRepository
}

func NewConsumerUsecase(userRepo models.UserRepository) models.ConsumerUsecase {
	return &consumerUsecase{userRepo}
}

func (u consumerUsecase) UserCreated(event events.UserCreatedEvent) error {
	_, err := u.userRepo.CreateUser(event.Name, event.Email)
	if err != nil {
		logs.Error(err)
		return errors.New("Can't create user")
	}
	return nil
}
func (u consumerUsecase) UserUpdated(event events.UserUpdatedEvent) error {
	_, err := u.userRepo.UpdateUser(event.ID, event.Name, event.Email)
	if err != nil {
		logs.Error(err)
		return errors.New("Can't update user")
	}
	return nil
}
func (u consumerUsecase) UserDeleted(event events.UserDeletedEvent) error {
	_, err := u.userRepo.DeleteUser(event.ID)
	if err != nil {
		logs.Error(err)
		return errors.New("failed to delete user")
	}
	return nil
}
