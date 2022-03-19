package provider

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

// SQLite is a provider that uses SQLite as the underlying database.
type SQLite struct {
	Database *sql.DB
}

// NewSQLite creates a new SQLite provider and opens the database. If the economy database does not exist, it will be created.
func NewSQLite(path string) (*SQLite, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(UUID TEXT NOT NULL, money UNSIGNED BIG INT NOT NULL);")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS economy_UUID_index ON economy(UUID);`)
	return &SQLite{db}, nil
}

// Balance ...
func (S *SQLite) Balance(UUID string) (uint64, error) {
	r := S.Database.QueryRow("SELECT money FROM economy WHERE UUID=$1", UUID)
	var money uint64
	err := r.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

// Set ...
func (S *SQLite) Set(UUID string, value uint64) error {
	_, err := S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", UUID, value)
	if err != nil {
		return err
	}
	return nil
}

// Decrease ...
func (S *SQLite) Decrease(UUID string, value uint64) error {
	bal, err := S.Balance(UUID)
	if err != nil {
		return err
	}
	_, err = S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", UUID, bal-value)
	if err != nil {
		return err
	}
	return nil
}

// Increase ...
func (S *SQLite) Increase(UUID string, value uint64) error {
	bal, err := S.Balance(UUID)
	if err != nil {
		return err
	}
	_, err = S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", UUID, bal+value)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the opened database connection and saves the sqlite file.
func (S *SQLite) Close() error {
	return S.Database.Close()
}
