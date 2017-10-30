package store

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	tmpDir string
)

func TestMain(main *testing.M) {
	tmpDir = fmt.Sprintf("./data_%x/", time.Now().UnixNano())
	if err := os.MkdirAll(tmpDir, os.ModeDir|os.ModePerm); err != nil {
		panic(err.Error())
	}

	main.Run()

	if err := os.RemoveAll(tmpDir); err != nil {
		panic(err.Error())
	}
}

func TestFunction(t *testing.T) {
	db := NewDB()
	if err := db.Open(tmpDir, tmpDir); err != nil {
		t.Fatalf("open db error:%v", err)
	}

	if err := db.Put("testkey", "testvalue"); err != nil {
		t.Fatalf("put value error:%v", err)
	}

	val, err := db.Get("testkey")
	if err != nil {
		t.Fatalf("get value error:%v", err)
	}
	t.Logf("get value:%s", val)
	db.Close()
}
