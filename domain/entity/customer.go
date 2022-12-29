package entity

import (
	"errors"
	"time"
)

type Customers struct {
	Id          int
	Name        string
	Address     string
	Phonenumber int
	Email       string
	Created_at  time.Time
	Update_at   time.Time
	Delete_at   time.Time
}

type DTOCustomers struct {
	Id          int
	Name        string
	Address     string
	PhoneNumber int
	Email       string
	Created_At  time.Time
	Update_At   time.Time
	Delete_At   time.Time
}

func NewCustomers(dto DTOCustomers) (*Customers, error) {
	if dto.Name == "" {
		return nil, errors.New("Name not be empty")
	}
	if dto.Address == "" {
		return nil, errors.New("Address not be empty")
	}
	if dto.PhoneNumber == 0 {
		return nil, errors.New("Phone Number not be empty")
	}
	if dto.Email == "" {
		return nil, errors.New("Email not be empty")
	}

	// strCreatedAt, _ := time.Parse("2022-01-02", dto.Created_At)
	// strUpdateAt, _ := time.Parse("2022-02-03", dto.Update_At)
	// strDeleteAt, _ := time.Parse("2022-03-04", dto.Delete_At)

	costumer := &Customers{
		Id:          dto.Id,
		Name:        dto.Name,
		Address:     dto.Address,
		Phonenumber: dto.PhoneNumber,
		Email:       dto.Email,
		Created_at:  dto.Created_At,
		Update_at:   dto.Update_At,
		Delete_at:   dto.Delete_At,
	}
	return costumer, nil
}

// Method Getter
func (c *Customers) GetIdCustomer() int {
	return c.Id
}

func (c *Customers) GetName() string {
	return c.Name
}
func (c *Customers) GetAddress() string {
	return c.Address
}
func (c *Customers) GetPhoneNumber() int {
	return c.Phonenumber
}
func (c *Customers) GetEmail() string {
	return c.Email
}
func (c *Customers) GetCreatedAt() string {
	return c.Created_at.Format("2022-01-02")
}

func (c *Customers) GetUpdateAt() string {
	return c.Update_at.Format("2022-02-03")
}
func (c *Customers) GetDeleteAt() string {
	return c.Delete_at.Format("2022-03-04")
}


//Method Setter 

func (cs *Customers) SetId(value int64) {
	cs.Id = int(value)
}

func (cs *Customers) SetName(value string) {
	cs.Name = value
}

func (cs *Customers) SetAddress(value string) {
	cs.Address = value
}

func (cs *Customers) SetPhoneNumber(value int64) {
	cs.Phonenumber = int(value)
}

func (cs *Customers) SetEmail(value string) {
	cs.Email = value
}

func (cs *Customers) SetCreatedAt(value time.Time) {
	cs.Created_at = value
}

func (cs *Customers) SetUpdateAt(value time.Time) {
	cs.Update_at = value
}

func (cs *Customers) SetDeleteAt(value time.Time) {
	cs.Delete_at = value
}
