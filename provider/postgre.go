package provider

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Postgre is a provider that uses PostgreSQL as the underlying database.
type Postgre struct {
	Database *sql.DB
	ConnectionDetails
}

type ConnectionDetails struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewPostgre creates a new Postgre provider and opens the database. If the economy database does not exist, it will be created.
func NewPostgre(c ConnectionDetails) (*SQLite, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", c.Host, c.Port, c.User, c.Password, c.Database))
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(UUID TEXT, money NUMERIC);")
	if err != nil {
		return nil, err
	}
	return &SQLite{db}, nil
}

// Balance ...
func (p *Postgre) Balance(UUID string) (uint64, error) {
	r := p.Database.QueryRow("SELECT money FROM economy WHERE UUID=?", UUID)
	var money uint64
	err := r.Scan(&money)
	if err != nil {
		return 0, err
	}
	return money, nil
}

// Set ...
func (p *Postgre) Set(UUID string, value uint64) error {
	_, err := p.Database.Exec("UPDATE economy SET money=$1 WHERE UUID=$2", value, UUID)
	if err != nil {
		return err
	}
	return nil
}

// Decrease ...
func (p *Postgre) Decrease(UUID string, value uint64) error {
	bal, err := p.Balance(XUID)
	if err != nil {
		return err
	}
	_, err = p.Database.Exec("UPDATE economy SET money=$1 WHERE UUID=$2", bal-value, UUID)
	if err != nil {
		return err
	}
	return nil
}

// Increase ...
func (p *Postgre) Increase(UUID string, value uint64) error {
	bal, err := p.Balance(UUID)
	if err != nil {
		return err
	}
	_, err = p.Database.Exec("UPDATE economy SET money=$1 WHERE XUID=$2", bal+value, UUID)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the opened database connection and saves the sqlite file.
func (p *Postgre) Close() error {
	return p.Database.Close()
}
