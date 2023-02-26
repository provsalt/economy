package provider

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Postgre is a provider that uses PostgreSQL as the underlying database.
type Postgre struct {
	database *sql.DB
}

type ConnectionDetails struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewPostgre creates a new Postgre provider with the given connection details.
func NewPostgre(details ConnectionDetails) (*Postgre, error) {
	connDet := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", details.Host, details.Port, details.User, details.Password, details.Database)
	db, err := sql.Open("postgres", connDet)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(UUID TEXT NOT NULL, balance BIGINT NOT NULL);")
	if err != nil {
		return nil, err
	}
	return &Postgre{database: db}, nil
}

// Balance ...
func (p Postgre) Balance(UUID string) (bal uint64, err error) {
	err = p.database.QueryRow("SELECT balance FROM economy WHERE UUID = $1", UUID).Scan(&bal)
	if err != nil {
		return 0, err
	}
	return
}

// Set ...
func (p Postgre) Set(UUID string, value uint64) (err error) {
	_, err = p.database.Exec("INSERT INTO economy VALUES ($1, $2) ON CONFLICT (UUID) DO UPDATE SET balance = $2", UUID, value)
	return err
}

// Close ...
func (p Postgre) Close() error {
	return p.database.Close()
}
