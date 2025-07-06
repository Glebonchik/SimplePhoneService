package db

import (
	"fmt"

	"danek.com/telephone/internal/domain"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewRepostitory(db *gorm.DB) domain.UserDatabase {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) AddUser(user *domain.User) error {
	result := r.db.Create(&user)
	if result.Error != nil {
		return domain.FormatErr("Adding new user", result.Error)
	}
	return nil
}

func (r *gormUserRepository) FindUser(id uint64) (domain.User, error) {

	var user domain.User
	result := r.db.First(&user, id)

	if result.Error != nil {
		return user, domain.FormatErr("Finding user", result.Error)
	}
	return user, nil

}

func (r *gormUserRepository) FindUserByPhone(phoneNum string) (domain.User, error) {
	var user domain.User

	result := r.db.Where("phone_number = ?", phoneNum).First(&user)
	if result.Error != nil {
		return user, domain.FormatErr("Finding user by phone", result.Error)
	}
	return user, nil
}

func (r *gormUserRepository) DeleteUser(id uint64) error {
	result := r.db.Delete(&domain.User{}, id)
	if result.Error != nil {
		return domain.FormatErr("Deliting user", result.Error)
	}
	return nil
}

func (r *gormUserRepository) UpdateUser(id uint64, user *domain.User) error {
	var userOnUpdate domain.User

	result := r.db.First(&userOnUpdate, id)
	if result.Error != nil {
		return domain.FormatErr("Updating User", result.Error)
	}

	userOnUpdate.Age = user.Age
	userOnUpdate.Contacts = user.Contacts
	userOnUpdate.Name = user.Name

	if err := r.db.Save(&userOnUpdate).Error; err != nil {
		return domain.FormatErr("Updating User", result.Error)
	}

	return nil
}

func (r *gormUserRepository) Print() error {
	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return domain.FormatErr("All user print", err)
	}
	fmt.Println(users)
	return nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(domain.User{})
}
