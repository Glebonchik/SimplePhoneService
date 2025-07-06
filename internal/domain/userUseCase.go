package domain

type UserUseCase interface {
	Register(user *User) error
	ClearUser(phone string) error
	UpdateUser(id uint64, user *User) error
	PrintUsers()
}
