package web

import "time"

type CustomerCreateRequest struct {
	Name        string    `validate:"required,max=200,min=1" json:"name"`
	Address     string    `validate:"required,max=200,min=1" json:"address"`
	PhoneNumber int       `validate:"required,max=100,min=1" json:"phone_number"`
	Email       string    `validate:"required,max=200,min=1" json:"email"`
	Created_at  time.Time `validate:"required, max=200,min=100 json:"created_at"`
	Update_at   time.Time `validate:"required, max=200,min=100 json:"created_at"`
	Delete_at   time.Time `validate:"required, max=200,min=100 json:"created_at"`
}
