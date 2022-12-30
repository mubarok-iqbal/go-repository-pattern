package repository

import "database/sql"

func CustomerRepository(db *sql.DB) CustomerRepositoryInterface {
	return &customerRepositoryMysql{DB: db}
}
