package repository

import (
	"context"

	"github.com/AndiVS/game_galaxy/internal/model"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// InsertUser function for inserting item from a table
func (repos *Postgres) InsertUser(c context.Context, user *model.User) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO users (id, account_login, universe_name, username) VALUES ($1, $2, $3) RETURNING id",
		user.ID, user.AccountLogin, user.UniverseName, user.Username)

	err := row.Scan(&user.ID)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectUser function for selecting item from a table
func (repos *Postgres) SelectUser(c context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User
	row := repos.Pool.QueryRow(c,
		"SELECT id, account_login, universe_name, username FROM users WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.AccountLogin, &user.UniverseName, &user.Username)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &user, err
	}

	log.Printf("sec")

	return &user, err
}

// DeleteUser function for deleting item from a table
func (repos *Postgres) DeleteUser(c context.Context, id uuid.UUID) error {
	ct, err := repos.Pool.Exec(c, "DELETE FROM users WHERE id = $1", id)

	if err != nil {
		return err
	} else if ct.RowsAffected() == 0 {
		log.Errorf("Not found : %s\n", err)
		return ErrNotFound
	}

	return nil
}
