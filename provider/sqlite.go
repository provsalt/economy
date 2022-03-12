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
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(XUID TEXT NOT NULL, money UNSIGNED BIG INT NOT NULL);")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS economy_xuid_index ON economy(XUID);`)
	return &SQLite{db}, nil
}

// Balance ...
func (S *SQLite) Balance(XUID string) (uint64, error) {
	r := S.Database.QueryRow("SELECT money FROM economy WHERE XUID=$1", XUID)
	var money uint64
	err := r.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

// Set ...
func (S *SQLite) Set(XUID string, value uint64) error {
	_, err := S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", XUID, value)
	if err != nil {
		return err
	}
	return nil
}

// Reduce ...
func (S *SQLite) Reduce(XUID string, value uint64) error {
	bal, err := S.Balance(XUID)
	if err != nil {
		return err
	}
	_, err = S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", XUID, bal-value)
	if err != nil {
		return err
	}
	return nil
}

// Increase ...
func (S *SQLite) Increase(XUID string, value uint64) error {
	bal, err := S.Balance(XUID)
	if err != nil {
		return err
	}
	_, err = S.Database.Exec("REPLACE INTO economy VALUES ($1, $2)", XUID, bal+value)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the opened database connection and saves the sqlite file.
func (S *SQLite) Close() error {
	return S.Database.Close()
}
