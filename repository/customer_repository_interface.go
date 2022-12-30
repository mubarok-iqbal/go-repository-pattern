package repository

import (
	"context"
	"go_repository_pattern/entity"
)

type CustomerRepositoryInterface interface {
	Insert(ctx context.Context, customer entity.Customer) (entity.Customer, error)
}