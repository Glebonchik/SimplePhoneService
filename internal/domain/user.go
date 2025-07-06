package domain

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Name     string
	Age      uint8
	Contacts UserContacts `gorm:"embedded"`
}

type UserContacts struct {
	PhoneNumber string
	Email       string
	Telegram    string
}
