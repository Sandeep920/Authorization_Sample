package services

import (
	"fmt"
	"go-graphql-jwt/graph/injection"
	"go-graphql-jwt/graph/model"
	"log"
)

// CreateUser 
func CreateUser(user *model.User) error {
	db := injection.DB

	var userData model.User

	userData.Name = user.Name
	userData.Email = user.Email
	userData.Password = user.Password
	userData.CreatedAt = user.CreatedAt
	userData.UpdatedAt = user.UpdatedAt

	if tx := db.Create(&user); tx.Error != nil {
		log.Print(tx.Error)
		return fmt.Errorf("error saving user")
	}

	return nil
}

// FindUsers 
func FindUsers() ([]*model.User, error) {
	db := injection.DB
	var users []*model.User

	if tx := db.Find(&users); tx.Error != nil {
		log.Print(tx.Error)
		return nil, tx.Error
	}
	return users, nil
}

// FindUserByEmail 
func FindUserByEmail(email string) (*model.User, error) {
	db := injection.DB
	var user *model.User

	if tx := db.Find(&user).Where("email = ?", email); tx.Error != nil {
		log.Print(tx.Error)
		return nil, tx.Error
	}
	return user, nil
}
