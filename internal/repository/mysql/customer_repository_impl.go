package repository

import (
	"Assigment-Restfull_Customer/domain/entity"
	"Assigment-Restfull_Customer/domain/repository"
	"Assigment-Restfull_Customer/internal/helper"
	"context"
	"database/sql"
	"errors"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() repository.CustomerRepository {
	return &CustomerRepositoryImpl{}
}

// Update implements repository.CustomerRepository
func (CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customers entity.Customers) entity.Customers {
	SQL := "insert into customer(name, address, phone, email, created_at, update_at, delete_at) values(?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, customers.Id, customers.Name, customers.Address, customers.Phonenumber, customers.Email, customers.Created_at, customers.Update_at, customers.Delete_at)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()

	helper.PanicIfError(err)

	customers.Id = int(id)

	return customers
}

// Update implements repository.CustomerRepository
func (CustomerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customers entity.Customers) entity.Customers {
	SQL := "update customer set name = ?, address = ?, phone = ?, email = ?, created_at = ?, update_at = ?, delete_at = ? where = ?"
	_, err := tx.ExecContext(ctx, SQL, customers.Name, customers.Address, customers.Phonenumber, customers.Email, customers.Created_at, customers.Update_at, customers.Delete_at, customers.Id)
	helper.PanicIfError(err)

	return customers
}

// Delete implements repository.CustomerRepository
func (CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customers entity.Customers) {
	SQL := "delete from customer where id = ?"
	_, err := tx.ExecContext(ctx, SQL, customers.Id)
	helper.PanicIfError(err)
}

// FindById implements repository.CustomerRepository
func (CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, costumersId int) (entity.Customers, error) {
	SQL := "select id, name , address, phone, email, created_at, update_at, delete_at from customers where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, costumersId)
	helper.PanicIfError(err)
	defer rows.Close()

	customers := entity.Customers{}
	if rows.Next() {
		err := rows.Scan(&customers.Id, &customers.Name, &customers.Address, customers.Phonenumber, customers.Email, customers.Created_at, customers.Update_at, customers.Delete_at)
		helper.PanicIfError(err)
		return customers, nil
	} else {
		return customers, errors.New("DATA NOT FOUND")
	}

}

// FindAll implements repository.CustomerRepository
func (CustomerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customers {
	SQL := "select id, name, address, phone, email, created_at, update_at, delete_at from customer"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customerss []entity.Customers
	for rows.Next() {
		customers := entity.Customers{}
		err := rows.Scan(&customers.Id, &customers.Name, &customers.Address, customers.Phonenumber, customers.Email, customers.Created_at, customers.Update_at, customers.Delete_at)
		helper.PanicIfError(err)
		customerss = append(customerss, customers)
	}
	return customerss
}
