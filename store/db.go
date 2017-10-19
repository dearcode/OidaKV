package store

import (
	"github.com/dgraph-io/badger"
)

type DB struct {
	badgerDB *badger.DB
}

func NewDB() *DB {
	db := new(DB)
	return db
}

func (db *DB) Open(keyPaht, valuePath string) error {
	opts := badger.DefaultOptions
	opts.Dir = keyPaht
	opts.ValueDir = valuePath
	badgerDB, err := badger.Open(opts)
	if err != nil {
		return err
	}
	db.badgerDB = badgerDB
	return nil
}

func (db *DB) Close() {
	if db.badgerDB != nil {
		db.badgerDB.Close()
		db.badgerDB = nil
	}
}

func (db *DB) Get(key string) ([]byte, error) {
	var val []byte
	err := db.badgerDB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		val, err = item.Value()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return val, err
}

func (db *DB) Put(key, val string) error {
	err := db.badgerDB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(val), 0)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
