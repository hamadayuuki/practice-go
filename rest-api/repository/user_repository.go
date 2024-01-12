package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
//                 引数                            戻り値
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// MARK: - User Repository の実体

type userRepository struct {
	db *gorm.DB
}

// DBからのDI用
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := userRepository.db.Where("email=?", email).First(user).Error; err != nil {   // TODO: Where() は今後実装
		return err
	}
	return nil
}

func (userRepository *userRepository) CreateUser(user *model.User) error {
	if err := userRepository.db.Create(user).Error; err != nil {   // TODO: Create() は今後実装
		return err
	}
	return nil
}