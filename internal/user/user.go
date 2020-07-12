package user

import (
	"context"
	"github.com/kopdar/kopdar-backend/pkg/errors"
)

var (
	ErrOrderNotFound = errors.E(errors.RequestFailed, errors.CodeNotFound, "user not found")
)

type Model struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Pin         string `json:"pin" db:"pin"`
}

type Service interface {
	FindAll(ctx context.Context) ([]*Model, error)
}

type Repository interface {
	FindAll(ctx context.Context) ([]*Model, error)
}

type service struct {
	repo Repository
}

// NewUserService creates a new user service.
func NewUserService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) FindAll(ctx context.Context) ([]*Model, error) {
	return s.repo.FindAll(ctx)
}
