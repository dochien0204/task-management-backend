package repository

import (
	"errors"
	"fmt"
	"log"
	"source-base-go/entity"
	"source-base-go/infrastructure/repository/define"
	"source-base-go/infrastructure/repository/util"

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

func (r UserRepository) GetListUser(statusId, page, size int, sortType, sortBy string) ([]*entity.User, error) {
	listUser := []*entity.User{}
	offset := util.CalculateOffset(page, size)
	if sortType == "" && sortBy == "" {
		sortType = "DESC"
		sortBy = "created_at"
	}
	if sortType == "" {
		sortType = "DESC"
	}

	switch sortBy {
	case "createdAt":
		sortBy = "created_at"

	case "updatedAt":
		sortBy = "updated_at"

	default:
		sortBy = "created_at"
	}

	err := r.db.Model(&entity.User{}).
		Where("status_id = ?", statusId).
		Offset(offset).
		Limit(size).
		Order(fmt.Sprintf("%v %v", sortBy, sortType)).
		Find(&listUser).Error
	
	if err != nil {
		return nil, err
	}

	return listUser, nil
}

func (r UserRepository) CountListUser(statusId int) (int, error) {
	var count int64
	err := r.db.Model(&entity.User{}).
		Where("status_id = ?", statusId).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r UserRepository) UpdateAvatar(userId int, avatar string) error {
	err := r.db.Model(&entity.User{}).
		Where("id = ?", userId).
		Update("avatar", avatar).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) ChangePassword(userId int , password string) error {
	err := r.db.Model(&entity.User{}).
		Where("id = ?", userId).
		Update("password", password).Error

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) IsUserEmailExists(email string) (bool, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

