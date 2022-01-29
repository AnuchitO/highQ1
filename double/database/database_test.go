package database

import (
	"testing"
)

type stub struct {
	Database
}

func (s *stub) Insert(collection string, data interface{}) error {
	return nil
}

func (s *stub) DBName() string {
	return "test db"
}

// type stubDB struct{}

// func (s *stubDB) Insert(collection string, data interface{}) error {
// 	return nil
// }

func TestInsert(t *testing.T) {
	s := &stub{}

	err := Insert(s, "product", `{}`)

	if err != nil {
		t.Error(err.Error())
	}
}
