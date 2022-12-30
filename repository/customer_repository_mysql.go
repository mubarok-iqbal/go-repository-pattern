package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_repository_pattern/entity"
	"strconv"
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

func (repository *customerRepositoryMysql) FindById(ctx context.Context, id uint32) (entity.Customer, error) {
	script := "SELECT id, name , email FROM customers WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	customer := entity.Customer{}
	if err != nil {
		return customer, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&customer.Id, &customer.Name, &customer.Email)
		return customer, nil
	} else {
		return customer, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
