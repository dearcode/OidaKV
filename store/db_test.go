package store

import (
	"testing"
)

func TestFunction(t *testing.T) {
	db := NewDB()
	err := db.Open("./data", "./data")
	if err != nil {
		t.Error("open db error %s", err)
	}

	err = db.Put("testkey", "testvalue")
	if err != nil {
		t.Error("put value error %s", err)
	}

	val, err := db.Get("testkey")
	if err != nil {
		t.Error("get value error %s", err)
	}
	t.Log("get value:%s", val)
	db.Close()
}
