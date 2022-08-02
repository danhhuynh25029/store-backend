package usecase

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"store/services/models"
	"store/services/repository"
)

type UserUsecase interface {
	AddUser(models.User) error
	GetUser(string) (*models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(useRepo repository.UserRepository) UserUsecase {
	return &userUsecase{useRepo}
}

func (u *userUsecase) AddUser(user models.User) error {
	return u.userRepo.AddUser(user)
}
func (u *userUsecase) GetUser(id string) (*models.User, error) {
	user_id, _ := primitive.ObjectIDFromHex(id)
	return u.userRepo.GetUser(user_id)
}
