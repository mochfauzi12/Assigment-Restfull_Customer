package web

import "time"

type CustomerUpdateRequest struct {
	Id          int       `validate:"required" json:"id"`
	Name        string    `validate:"required,max=200,min=1" json:"name"`
	Address     string    `validate:"required,max=200,min=1" json:"address"`
	PhoneNumber int       `validate:"required,max=100,min=1" json:"phone_number"`
	Email       string    `validate:"required,max=200,min=1" json:"email"`
	Created_at  time.Time `validate:"required,  json:"created_at"`
	Update_at   time.Time `validate:"required, json:"created_at"`
	Delete_at   time.Time `validate:"required,  json:"created_at"`
}
