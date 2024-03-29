package go_repository_pattern

import (
	"context"
	"database/sql"
	"fmt"
	"go_repository_pattern/entity"
	"go_repository_pattern/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/golang_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func TestCustomerInsert(t *testing.T) {
	customerRepository := repository.CustomerRepository(GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Name:  "Faqih Muhammad",
		Email: "faqih@test.com",
	}

	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCustomerFindById(t *testing.T) {
	customerRepository := repository.CustomerRepository(GetConnection())

	customer, err := customerRepository.FindById(context.Background(), 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestCustomerFindAll(t *testing.T) {
	customerRepository := repository.CustomerRepository(GetConnection())

	customers, err := customerRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}
}
