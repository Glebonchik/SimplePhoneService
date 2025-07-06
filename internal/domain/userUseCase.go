package domain

type UserUseCase interface {
	Register(user *User) error
	ClearUser(phone string) error
	UpdateUser(id uint64, user *User) error
	FindUserByPhone(phone string) (User, error)
	PrintUsers()
}
