package usecase

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2/log"

	"service1/logs"
	"service1/modules/entities/models"
)

type userService struct {
	userRepo models.UserRepository
	producer models.ProducerUser
}

func NewUserService(userRepo models.UserRepository, producer models.ProducerUser) models.UserUsecase {
	return &userService{
		userRepo: userRepo,
		producer: producer,
	}
}

func (u userService) Register(name string, email string) (*models.UserRequest, error) {
	newUser, err := u.userRepo.CreateUser(name, email)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("Can't create user")
	}
	message := &models.UserRequest{
		Id:    uint(newUser.ID),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	if err = u.producer.UserCreated(message, time.Now()); err != nil {
		log.Error(err)
	}

	log.Info("create user Successfuly")
	return message, nil
}

func (u userService) UpdateAccount(id uint, req models.UserRequest) (*models.UserRequest, error) {

	updateUser, err := u.userRepo.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	message := &models.UserRequest{
		Id:    id,
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}
	err = u.producer.UserUpdated(message, time.Now())
	if err != nil {
		log.Error(err)
	}
	log.Info("update user Successfuly")
	return message, nil
}

func (u userService) DeleteAccount(id uint) (*string, error) {
	deleteMessage, err := u.userRepo.DeleteUser(id)
	if err != nil {
		logs.Error(err)
		return nil, errors.New("failed to delete user")
	}
	if err := u.producer.UserDeleted(id); err != nil {
		log.Error(err)
		return nil, errors.New("failed to delete user")
	}
	log.Info(deleteMessage)
	return deleteMessage, nil
}
