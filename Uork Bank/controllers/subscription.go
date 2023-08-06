package controllers

import (
	"uorkbank/db"
)

type UserSubscription struct {
	Nome  string
	Email string
	Senha string
	Saldo int
}

func CreateUserSubscription(nome, email, senha string, saldo int) error {

	s := UserSubscription{Nome: nome, Email: email, Senha: senha, Saldo: saldo}

	return db.Insert("uork_users", s)
}



