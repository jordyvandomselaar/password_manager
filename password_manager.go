package main

import (
	"github.com/jordyvandomselaar/password_manager/actions"
	"github.com/dgraph-io/badger"
	"github.com/jordyvandomselaar/password_manager/database"
	password2 "github.com/jordyvandomselaar/password_manager/password"
	"os"
)

func main() {
	options := badger.DefaultOptions
	options.Dir = "./db/keys"
	options.ValueDir = "./db/values"

	db := database.Open(options)
	defer db.Close()

	action := os.Args[1]

	if action == "save" {
		save(db)
	}
}

func save(db *badger.DB) {
	name, username, password, description := actions.GetArgs()

	passwordStruct := password2.Password{
		Name:        name,
		Username:    username,
		Password:    password,
		Description: description,
	}

	password2.Save(db, passwordStruct)
}
