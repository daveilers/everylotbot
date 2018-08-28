package db

import (
	"database/sql"
	"fmt"

	"github.com/daveilers/everylotbot/everylot"
	_ "github.com/mattn/go-sqlite3"
)

type Conn struct {
	db *sql.DB
}

func NewSQLite(DBDriver, DBDSN string) (db *Conn, err error) {
	c, err := sql.Open(DBDriver, DBDSN)
	if err != nil {
		return nil, err
	}
	return &Conn{db: c}, nil
}

func (c *Conn) MarkAsTweeted(l *everylot.Lot, statusID string) (err error) {
	_, err = c.db.Exec("UPDATE lots SET tweeted = ? WHERE id = ?", statusID, l.ID)
	return
}

func (c *Conn) Next() (l *everylot.Lot, err error) {

	var lat, lon sql.NullFloat64
	row := c.db.QueryRow("SELECT id, lat, lon, address, city, state, zip, tweeted FROM lots WHERE tweeted = 0 ORDER BY id LIMIT 1;")
	if row == nil {
		return nil, fmt.Errorf("No results")
	}
	l = &everylot.Lot{}
	err = row.Scan(&l.ID, &lat, &lon, &l.Address, &l.City, &l.State, &l.Zip, &l.Tweeted)
	if lat.Valid && lon.Valid {
		l.Lat = lat.Float64
		l.Lon = lon.Float64
	}

	return
}

func (c *Conn) ID(id string) (l *everylot.Lot, err error) {

	row := c.db.QueryRow(`SELECT id, lat, lon, address, city, state, zip, tweeted 
	FROM lots 
	WHERE tweeted = 0 
	AND id = ?
	ORDER BY id 
	LIMIT 1;`, id)
	if row == nil {
		return nil, fmt.Errorf("No results")
	}
	l = &everylot.Lot{}
	err = row.Scan(&l.ID, &l.Lat, &l.Lon, &l.Address, &l.City, &l.State, &l.Zip, &l.Tweeted)
	return
}
