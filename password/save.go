package password

import (
	"github.com/dgraph-io/badger"
)

func Save(db *badger.DB, password Password) {
	db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(password.Name), Marshal(password), 0)

		if err != nil {
			panic(err)
		}

		return nil
	})
}
