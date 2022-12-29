package usecase

import (
	"Assigment-Restfull_Customer/domain/entity"
	"Assigment-Restfull_Customer/domain/repository"
	"Assigment-Restfull_Customer/domain/usecase"
	"Assigment-Restfull_Customer/domain/web"
	"Assigment-Restfull_Customer/exception"
	"Assigment-Restfull_Customer/internal/helper"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type CustomerUseCaseImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository, DB *sql.DB, validated *validator.Validate) usecase.CustomerUseCase {
	return &CustomerUseCaseImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validated,
	}
}

// Create implements usecase.CustomerUseCase
func (service *CustomerUseCaseImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	customer := entity.Customers{
		Name:        request.Name,
		Address:     request.Address,
		Phonenumber: request.PhoneNumber,
		Email:       request.Email,
		// Created_At:  request.Created_at,
		// Update_At:   request.Update_at,
		// Delete_At:   request.Delete_at,
	}

	customer = service.CustomerRepository.Save(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)

}

// Update implements usecase.CustomerUseCase
func (service *CustomerUseCaseImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customery, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customery.Name = request.Name
	customery.Address = request.Address
	customery.Phonenumber = request.PhoneNumber
	customery.Email = request.Email
	// customers.Created_At = request.Created_at
	// customers.Update_At = request.Update_at
	// customers.Delete_At = request.Delete_at

	service.CustomerRepository.Update(ctx, tx, customery)

	//return helper.ToCustomerResponse(customers)
	return helper.ToCustomerResponse(customery)

}

// Delete implements usecase.CustomerUseCase
func (service *CustomerUseCaseImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customery, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customery)

}

// FindById implements usecase.CustomerUseCase
func (service *CustomerUseCaseImpl) FindById(ctx context.Context, customerId int) web.CustomerResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customery, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customery)

}

// FindAll implements usecase.CustomerUseCase
func (service *CustomerUseCaseImpl) FindAll(ctx context.Context) []web.CustomerResponse {

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customery := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(customery)

}
