package utils

import (
	"time"
	"user-registration/cmd/initializers"
	"user-registration/cmd/model"
)

func SaveUser(user model.User) (error, *model.User) {
	u := model.User{
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		Password: user.Password,
		Address:  user.Address,
		Valid:    true,
		ValidAt:  time.Now(),
	}
	result := initializers.DB.Create(&u)
	if result.Error != nil {
		return result.Error, nil
	}
	return nil, &u
}
