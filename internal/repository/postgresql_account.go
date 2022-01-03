package repository

import (
	"context"
	"github.com/AndiVS/game_galaxy/internal/model"
	log "github.com/sirupsen/logrus"
)

// InsertAccount function for inserting item from a table
func (repos *Postgres) InsertAccount(c context.Context, account *model.Account) error {
	row := repos.Pool.QueryRow(c,
		"INSERT INTO accounts(login, password) VALUES ($1, $2) RETURNING login", account.Login, account.Password)

	err := row.Scan(&account.Login)
	if err != nil {
		log.Errorf("Unable to INSERT: %v", err)
		return err
	}

	return err
}

// SelectAccount function for selecting item from a table
func (repos *Postgres) SelectAccount(c context.Context, login string) (*model.Account, error) {
	var account model.Account
	row := repos.Pool.QueryRow(c,
		"SELECT login, password FROM accounts WHERE login = $1", login)

	err := row.Scan(&account.Login, &account.Password)
	if err != nil {
		log.Errorf("Unable to SELECT: %v", err)
		return &account, err
	}

	log.Printf("sec")

	return &account, err
}
