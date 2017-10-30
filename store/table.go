package store

import (
	"OidaKV/proto"
	"os"
)

type Table struct {
	db *DB
}

func NewTable(keyPath, valuePath string) *Table {
	os.MkdirAll(keyPath, os.ModeDir|os.ModePerm)
	os.MkdirAll(valuePath, os.ModeDir|os.ModePerm)

	table := new(Table)
	table.db = NewDB()
	err := table.db.Open(keyPath, valuePath)
	if err != nil {
		log.Println("Open DB failed: ", err)
		return nil
	}
	return table
}

func (table *Table) Close() {
	table.db.Close()
}

func (table *Table) Get(keyValue proto.KeyValue) {
}
