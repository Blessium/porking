package service

import (
	"errors"
	"github.com/blessium/porking/database"
	"github.com/blessium/porking/model"
)

type UserService struct {
}

func (u *UserService) GetUserById(id uint) (*model.User, error) {

	user := new(model.User)

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	result := db.Limit(1).First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (u *UserService) GetUsers() (*[]model.User, error) {
	users := new([]model.User)

	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	result := db.Find(users)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("User not found")
	}

	return users, nil
}

func (u *UserService) UpdateUser(info *model.UserInfo, id uint) error {
	user, err := u.GetUserById(id)
	if err != nil {
		return err
	}

	db, err := database.ConnectDatabase()

	if err != nil {
		return err
	}
	db.Save(info.UpdateUser(user))

	return nil
}

func (u *UserService) RegisterUser(user *model.User) error {
	db, err := database.ConnectDatabase()

	if err != nil {
		return err
	}

	db.Save(user)
	return nil
}

func (u *UserService) AuthUser(login *model.UserLogin) (*model.User, error) {
	db, err := database.ConnectDatabase()

	if err != nil {
		return nil, err
	}

	user := new(model.User)
	result := db.Limit(1).Where(&model.User{Email: login.Email, Password: login.Password}).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("Invalid password or email")
	}

	return user, nil
}
