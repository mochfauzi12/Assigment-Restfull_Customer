package helper

import (
	"Assigment-Restfull_Customer/domain/entity"
	"Assigment-Restfull_Customer/domain/web"
)

func ToCustomerResponse(customer entity.Customers) web.CustomerResponse {
	return web.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.Phonenumber,
		Email:       customer.Email,
		Created_At:  customer.Created_at,
		Update_At:   customer.Update_at,
		Delete_At:   customer.Delete_at,
	}
}

func ToCustomerResponses(customerss []entity.Customers) []web.CustomerResponse {
	var customerResponses []web.CustomerResponse
	for _, customer := range customerss {
		customerResponses = append(customerResponses, ToCustomerResponse(customer))
	}
	return customerResponses
}
