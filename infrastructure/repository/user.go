package repository

import (
	"errors"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepostory(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) WithTrx(trxHanlde *gorm.DB) UserRepository {
	if trxHanlde == nil {
		log.Print("Transaction DB not found")
		return r
	}

	r.db = trxHanlde
	return r
}

func (r UserRepository) GetUserProfile(userId int) (*entity.User, error) {
	user := entity.User{}

	result := r.db.Model(&entity.User{}).
		Where("id = ?", userId).
		Preload("Status").
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r UserRepository) FindByUsername(userName string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.
		Where("username = ?", userName).
		Preload("Status").
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (r UserRepository) CreateUser(user *entity.User) error {
	//Hash password
	err := user.HashPassword()
	if err != nil {
		return err
	}

	//Add status active for user
	statusActive := &entity.Status{}
	err = r.db.Model(&entity.Status{}).
		Where("code = ?", define.ACTIVE).
		First(&statusActive).Error
	if err != nil {
		return err
	}
	user.StatusId = statusActive.Id

	//Create user
	err = r.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) IsUserExists(username string) (bool, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}
