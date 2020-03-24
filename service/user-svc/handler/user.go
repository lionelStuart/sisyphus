package handler

import (
	"sisyphus/common/app"
	"sisyphus/models"
	"sisyphus/models/po"
)

type User struct {
	ID string
	// Uid      string
	Username string
	Password string
	Email    string
	Phone    string
	State    int

	Profile Profile

	PageNum  int
	PageSize int
}

type Profile struct {
	Nickname string
	Age      int8
	Gender   string
	Address  string
}

func (u *User) Add() error {
	user := app.H{
		"username": u.Username,
		"password": u.Password,
		"email":    u.Email,
		"phone":    u.Phone,
		"state":    u.State,
		"profile": app.H{
			"nickname": u.Profile.Nickname,
			"age":      u.Profile.Age,
			"gender":   u.Profile.Gender,
			"address":  u.Profile.Address,
		},
	}

	if err := models.AddAuthProfile(user); err != nil {
		return err
	}
	return nil
}

func (u *User) EditAuth() error {
	return models.EditAuth(u.ID, app.H{
		"username": u.Username,
		"password": u.Password,
		"email":    u.Email,
		"phone":    u.Phone,
		"state":    u.State,
	})
}

func (u *User) EditProfile() error {
	return models.EditProfile(u.ID, app.H{
		"nickname": u.Profile.Nickname,
		"age":      u.Profile.Age,
		"gender":   u.Profile.Gender,
		"address":  u.Profile.Address,
	})
}

func (u *User) GetProfile() (*po.Profile, error) {
	return models.GetProfileById(u.ID)
}
