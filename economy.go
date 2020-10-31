package Economy

import (
	"database/sql"
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/player"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Economy struct {
	server *dragonfly.Server

	path string

	config string
}

var Db *sql.DB

func New(s *dragonfly.Server, path string, config string) *Economy {
	e := &Economy{
		server: s,
		path:   path,
		config: config,
	}
	return e
}

func (e Economy) StartDB() (sql.Result, error) {
	db, err := sql.Open("mysql", e.config)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(30)
	Db = db
	res, err := db.Exec("CREATE TABLE IF NOT EXISTS economy ( XUID VARCHAR(36) PRIMARY KEY, username TEXT NOT NULL, money FLOAT DEFAULT 0);")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (e Economy) InitPlayer(player *player.Player, defaultmoney float64) (error, sql.Result) {
	r, err := Db.Query("SELECT username FROM economy WHERE XUID=?", player.XUID())
	if err != nil {
		return err, nil
	}
	if !r.Next() {
		res, err := Db.Exec("REPLACE INTO economy (XUID, username, money) VALUES (?, ?, ?)", player.XUID(), player.Name(), defaultmoney)
		if err != nil {
			return nil, res
		}
		return nil, res
	}
	return nil, nil
}
