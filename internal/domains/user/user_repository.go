package user

import (
	"context"
	"go-clean-architecture/internal/domains/user/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := u.database.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindUser(ctx context.Context, id string) (*entity.User, error) {
	var user *entity.User
	if err := u.database.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	if err := u.database.WithContext(ctx).Where("email = ?", email).First(&entity.User{}).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	if err := u.database.WithContext(ctx).Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := u.database.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := u.database.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserRepository(database *gorm.DB) IUserRepository {
	return &userRepository{
		database: database,
	}
}
