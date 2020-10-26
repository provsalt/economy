package Economy

import (
	"database/sql"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"time"
)

type Economy struct {
	server *dragonfly.Server

	path string

	config string
}

func New(s *dragonfly.Server, path string, config string) (*Economy, error) {
	e := &Economy{
		server: s,
		path:   path,
		config: config,
	}
	db, err := sql.Open("mysql", e.config)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
		"UUID VARCHAR(36) PRIMARY KEY," +
		"username TEXT NOT NULL," +
		"money FLOAT DEFAULT 0",
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e Economy) InitPlayer(player *player.Player, defaultmoney float64) error {
	panic("TODO")
}

func (e Economy) GetMoney(player *player.Player) error {
	panic("TODO")
}
