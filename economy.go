package economy

import (
	"database/sql"
	"errors"
	"github.com/df-mc/dragonfly/dragonfly/player"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Economy struct {
	Database *sql.DB
}

type Connection struct {
	IP       string
	Username string
	Password string
	Schema   string
}

func New(connection Connection, minConn int, maxconn int) Economy {
	db, err := sql.Open("mysql", connection.Username+":"+connection.Password+"@("+connection.IP+")/"+connection.Schema)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(maxconn)
	db.SetMaxIdleConns(minConn)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS economy(XUID BIGINT, username TEXT, money FLOAT);")
	if err != nil {
		panic(err)
	}
	return Economy{
		Database: db,
	}
}

func (e Economy) InitPlayer(player *player.Player, defaultmoney int) bool {
	r := e.Database.QueryRow("SELECT XUID FROM economy WHERE username=?", player.XUID())
	var XUID int
	err := r.Scan(&XUID)
	if err == nil {
		return true
	}
	if errors.Is(err, sql.ErrNoRows) {
		_, err := e.Database.Exec("REPLACE INTO economy (XUID, username, money) VALUES (?, ?, ?)", player.XUID(), player.Name(), defaultmoney)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
	return true
}
func (e Economy) Close() {
	e.Database.Close()
}

func (e Economy) Balance(player *player.Player) (error, int) {
	r := e.Database.QueryRow("SELECT money FROM economy WHERE XUID=?", player.XUID())
	var money int
	err := r.Scan(&money)
	if err != nil {
		return err, 0
	}
	return nil, money
}
