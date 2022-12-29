package web

import "time"

type CustomerResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber int       `json:"phone_number"`
	Email       string    `json:"email"`
	Created_At  time.Time `json:"created_at"`
	Update_At   time.Time `json:"update_at"`
	Delete_At   time.Time `json:"delete_at"`
}
