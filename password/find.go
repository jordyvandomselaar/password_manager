package password

import (
	"github.com/dgraph-io/badger"
)

func Find(db *badger.DB, name string) Password {
	var data []byte

	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(name))

		if err != nil {
			panic(err)
		}

		value, err := item.Value()

		if err != nil {
			panic(err)
		}

		data = make([]byte, len(value))

		copy(data, value)

		return nil
	})

	return UnMarshal(data)
}
