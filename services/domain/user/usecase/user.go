package usecase

import (
	"store/services/models"
	"store/services/repository/mgo"
)

type UserUsecase interface {
	AddUser(models.User) error
	GetUser(string) (*models.User, error)
}

type userUsecase struct {
	userRepo mgo.UserRepository
}

func NewUserUsecase(useRepo mgo.UserRepository) UserUsecase {
	return &userUsecase{useRepo}
}

func (u *userUsecase) AddUser(user models.User) error {
	return u.userRepo.AddUser(user)
}
func (u *userUsecase) GetUser(objId string) (*models.User, error) {

	return u.userRepo.GetUser(objId)
}
func (u *userUsecase) DeleteUser(objId string) error {
	return u.userRepo.DeleteUser(objId)
}
func (u *userUsecase) UpdateUser(obj string, user models.User) error {
	return u.userRepo.UpdateUser(obj, user)
}
