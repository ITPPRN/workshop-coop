package usecase

import (
	"github.com/gofiber/fiber/v2/log"

	"service1/modules/entities/events"
	"service1/modules/entities/models"
)

type ConsumerService interface {
	UserReaded(event events.UserReadedEvent) error
}

type consumerService struct {
	repo models.ConsumerRepository
}

func NewConsumerUsecase(comsumerRepo models.ConsumerRepository) ConsumerService {
	return &consumerService{repo: comsumerRepo}
}

func (u *consumerService) UserReaded(event events.UserReadedEvent) error {

	err := u.repo.CreateUserReadedDog(&models.UserReadDog{
		CreateAt:   event.TimeStamp,
		UserID:     event.UserId,
		DogID:      event.DogId,
		DogDetails: event.DogDetails,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil

}
