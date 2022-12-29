package repository

import (
	"Assigment-Restfull_Customer/domain/entity"
	"context"
	"database/sql"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customers entity.Customers) entity.Customers
	Update(ctx context.Context, tx *sql.Tx, customers entity.Customers) entity.Customers
	Delete(ctx context.Context, tx *sql.Tx, customers entity.Customers)
	FindById(ctx context.Context, tx *sql.Tx, costumersId int) (entity.Customers, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Customers
}
