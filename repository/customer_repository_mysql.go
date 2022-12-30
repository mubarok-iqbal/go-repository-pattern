package repository

import (
	"context"
	"database/sql"
	"go_repository_pattern/entity"
)

type customerRepositoryMysql struct {
	DB *sql.DB
}

func (repository *customerRepositoryMysql) Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := "INSERT INTO customers(name, email) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, customer.Name, customer.Email)
	if err != nil {
		return customer, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return customer, err
	}
	customer.Id = uint32(id)
	return customer, nil
}
