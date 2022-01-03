// Package service encapsulates
package service

import (
	"context"

	"github.com/AndiVS/game_galaxy/internal/repository"
	"github.com/google/uuid"
)

// Storage interface for mocks
type Storage interface {
	UpdateStorage(c context.Context, id uuid.UUID) error
}

// StorageService for generating token
type StorageService struct {
	Rep repository.Storage
}

// NewServiceUser  for setting new authorizer
func NewServiceUser(repositories interface{}) Storage {
	return &StorageService{Rep: repositories.(*repository.Postgres)}
}

// UpdateStorage function for deleting item from a table
func (s *StorageService) UpdateStorage(c context.Context, id uuid.UUID) error {
	return s.Rep.UpdateStorage(c, id)
}
