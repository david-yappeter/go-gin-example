package service

import (
	"fmt"
	"myapp/config"
	"myapp/entity"
	"myapp/tools"

	"github.com/google/uuid"
)

type UserService interface {
	Create(entity.User) (*entity.User, error)
	Update(entity.User) (*entity.User, error)
	Delete(string) error
	GetByID(uuID string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Login(entity.UserLoginForm) (jwtToken *string, err error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Create(input entity.User) (*entity.User, error) {
	db := config.DB()

	input.UuID = uuid.New().String()
	input.Password = tools.HashPassword(input.Password)
	if err := db.Model(input).Create(&input).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &input, nil
}

func (s *userService) Update(input entity.User) (*entity.User, error) {
	db := config.DB()

	if err := db.Model(input).Where("uu_id = ?", input.UuID).Updates(&input).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &input, nil
}

func (s *userService) Delete(uuID string) error {
	db := config.DB()

	if err := db.Model(&entity.User{}).Delete(&entity.User{}, uuID).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *userService) GetByID(uuID string) (*entity.User, error) {
	db := config.DB()

	var user entity.User
	if err := db.Model(entity.User{}).Where("uu_id  = ?", uuID).Take(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetByEmail(email string) (*entity.User, error) {
	db := config.DB()

	var user entity.User
	if err := db.Model(entity.User{}).Where("email  = ?", email).Take(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *userService) GetAll() ([]*entity.User, error) {
	db := config.DB()

	var users []*entity.User
	if err := db.Model(entity.User{}).Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

func (s *userService) Login(input entity.UserLoginForm) (*string, error) {
	userByID, err := s.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	if err := tools.CompareHash(userByID.Password, input.Password); err != nil {
		return nil, err
	}

	token := JwtGenerate(userByID.UuID)
	return &token, nil
}
