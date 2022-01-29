package sql

import (
	"database/sql"
	"testing"
)

type stub struct {
	query        string
	lastInsertID int64
	rowsAffected int64
}

func (m *stub) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}
func (m *stub) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *stub) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}

func TestExecQuery(t *testing.T) {
	s := &stub{rowsAffected: 1}

	r, err := execQuery(s, "SELECT 1", nil)
	if err != nil {
		t.Fatal(err)
	}

	if r != 1 {
		t.Fatalf("Expected 1, got %d", r)
	}
}
