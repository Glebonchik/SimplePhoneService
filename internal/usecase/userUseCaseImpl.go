package usecase

import (
	"danek.com/telephone/internal/domain"
)

type UserUseCase struct {
	repo domain.UserDatabase
}

func NewUserUseCase(repo domain.UserDatabase) domain.UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(user *domain.User) error {
	return uc.repo.AddUser(user)
}

func (uc *UserUseCase) ClearUser(phone string) error {
	var user domain.User
	user, err := uc.repo.FindUserByPhone(phone)
	if err != nil {
		return err
	}
	return uc.repo.DeleteUser(user.ID)
}

func (uc *UserUseCase) UpdateUser(id uint64, user *domain.User) error {
	return uc.repo.UpdateUser(id, user)
}

func (uc *UserUseCase) PrintUsers() {
	uc.repo.Print()
}
