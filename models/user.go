package models

import (
	"errors"
	"fmt"
)

// User structure
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers : List all users
func GetUsers() []*User {
	return users
}

// AddUsers : add a new user element
func AddUsers(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include ID, or it must be zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID : Retrieve a user by specifying the userId
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("Could not find user with ID = %v", id)
}

// UpdateUser : Update a user
func UpdateUser(u User) (User, error) {
	for i, existing := range users {
		if u.ID == existing.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("Could not find user with ID = %v", u.ID)
}

// RemoveUserByID : Remove a user with a little splicing fun
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Failed to find element with ID = %v", id)
}
