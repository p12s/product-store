package service

import (
	"github.com/p12s/product-store/server/internal/repository"
)

// Service
type Service struct {
	Producter
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Producter: NewProductService(repos.Producter),
	}
}
