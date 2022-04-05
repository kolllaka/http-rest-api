package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.QueryRow("SET FOREIGN_KEY_CHECKS = 0;")
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
			db.QueryRow("SET FOREIGN_KEY_CHECKS = 1;")
		}

		db.Close()
	}
}
