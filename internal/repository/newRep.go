package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	// ErrNotFound means entity is not found in repository
	ErrNotFound = errors.New("not found")
)

// Postgres struct for Pool
type Postgres struct {
	Pool *pgxpool.Pool
}

// NewRepository set new repository for mongo and postgres
func NewRepository(db interface{}) Account {
	return &Postgres{Pool: db.(*pgxpool.Pool)}
}

// Account used for structuring, function for working with accounts
type Account interface {
	InsertAccount(c context.Context, account *model.Account) error
	SelectAccount(c context.Context, login string) (*model.Account, error)
}

// Universe used for structuring, function for working with universes
type Universe interface {
	InsertUniverse(c context.Context, universe *model.Universe) error
	SelectUniverse(c context.Context, name string) (*model.Universe, error)
	DeleteUniverse(c context.Context, name string) error
}

// Galaxy used for structuring, function for working with galaxies
type Galaxy interface {
	InsertGalaxy(c context.Context, galaxy *model.Galaxy) error
	SelectGalaxy(c context.Context, id uuid.UUID) (*model.Galaxy, error)
	DeleteGalaxy(c context.Context, id uuid.UUID) error
}

// System used for structuring, function for working with galaxies
type System interface {
	InsertSystem(c context.Context, system *model.System) error
	SelectSystem(c context.Context, id uuid.UUID) (*model.System, error)
	DeleteSystem(c context.Context, id uuid.UUID) error
}

// Users used for structuring, function for working with users
type Users interface {
	InsertUser(c context.Context, user *model.User) error
	SelectUser(c context.Context, username string) (*model.User, error)
	SelectAllUser(c context.Context) ([]*model.User, error)
	UpdateUser(c context.Context, username string, isAdmin bool) error
	DeleteUser(c context.Context, username string) error
}
