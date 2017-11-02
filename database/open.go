package database

import (
	"github.com/dgraph-io/badger"
)

func Open(options badger.Options) *badger.DB {
	db, err := badger.Open(options)

	if err != nil {
		panic(err)
	}

	return db
}
