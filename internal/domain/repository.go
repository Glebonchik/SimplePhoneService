package domain

type UserDatabase interface {
	AddUser(user *User) (err error)
	FindUser(id uint64) (user User, err error)
	FindUserByPhone(phoneNum string) (user User, err error)
	DeleteUser(id uint64) (err error)
	UpdateUser(id uint64, userData *User) (err error)
	Print() error
}
