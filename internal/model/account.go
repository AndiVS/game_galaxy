// Package model all models that will be used in project
package model

type Account struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
