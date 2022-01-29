package sql

import (
	"database/sql"
	"testing"
)

func TestExecQuery(t *testing.T) {
	r, err := execQuery(&sql.DB{}, "SELECT 1", nil)
	if err != nil {
		t.Fatal(err)
	}

	if r != 1 {
		t.Fatalf("Expected 1, got %d", r)
	}
}
