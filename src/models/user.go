package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string       `json:"name" bson:"name"`
	Email        string       `json:"email" bson:"email" gorm:"uniqueIndex"`
	Password     string       `json:"password" bson:"password"`
	Age          int          `json:"age" bson:"age"`
	Country      string       `json:"country" bson:"country"`
	Profession   string       `json:"profession" bson:"profession"`
	Organization Organization `json:"organization" bson:"organization" gorm:"embedded"`
}

type Organization struct {
	gorm.Model
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

// Saves a user to the database
func (user *User) Save() (*User, error) {
	err := Database.Model(&user).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Fetches all users from the database
func FetchAllUsers() (*[]User, error) {
	var users []User
	err := Database.Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

// Fetches a user from the database
func FetchUser(id string) (*User, error) {
	var user User
	err := Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// Updates a user in the database
func (user *User) UpdateUser(id string) (*User, error) {
	err := Database.Model(&User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// Deletes a user from the database
func DeleteUser(id string) error {
	err := Database.Model(&User{}).Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return err
	}
	return nil
}
