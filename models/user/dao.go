package user

import (
	"errors"
	"gorm.io/gorm"
	"sign/models"
)

func NewIUser() IUser {
	return &User{
		db: models.GetDb(),
	}
}

type IUser interface {
	Create(in *models.User) error
	PhoneExist(phone string) (bool, error)
	GetByPhone(phone string) (*models.User, error)
	Get(id int64) (*models.User, error)
}

type User struct {
	db *gorm.DB
}

func (u *User) Get(id int64) (*models.User, error) {
	info := new(models.User)
	err := u.db.Where("id=?", id).First(info).Error
	return info, err
}

func (u *User) GetByPhone(phone string) (*models.User, error) {
	info := new(models.User)
	err := u.db.Where("phone=?", phone).First(info).Error
	return info, err
}

func (u *User) PhoneExist(phone string) (bool, error) {
	err := u.db.Where("phone=?", phone).First(new(models.User)).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return true, err
		}
	}
	return true, nil
}

func (u *User) Create(in *models.User) error {
	return u.db.Create(in).Error
}
