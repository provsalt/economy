package Economy

import (
	"database/sql"
	"github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Economy struct {
	server *dragonfly.Server

	path string

	config struct{
		User             string            // Username
		Passwd           string            // Password (requires User)
		Addr             string            // Network address (requires Net)
		DBName           string            // Database name
	}
}

func New(s *dragonfly.Server, path string, config struct{
	User             string            // Username
	Passwd           string            // Password (requires User)
	Addr             string            // Network address (requires Net)
	DBName           string            // Database name
}) (*Economy, error){
		e := &Economy{
		server: s,
		path:   path,
		config: config,
	}
		config2 := mysql.Config{
		User:   e.config.User,
		Passwd: e.config.Passwd,
		Addr:   e.config.Addr,
		DBName: e.config.DBName,
	}
		db, err := sql.Open("mysql", config2.FormatDSN())
		if err != nil {
		return nil, err
	}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(2)

		r, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
		"UUID VARCHAR(36) PRIMARY KEY," +
		"username TEXT NOT NULL," +
		"money FLOAT DEFAULT 0",
	)
		spew.Dump(r)
		if err != nil {
		return nil, err
	}
	return e, nil
}

func (e Economy) InitPlayer(player *player.Player, defaultmoney float64 ) error{
	panic("TODO")
}

func (e Economy) GetMoney(player *player.Player)  error{
	panic("TODO")
}
