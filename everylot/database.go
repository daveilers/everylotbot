package everylot

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func openIt() (db *sql.DB, err error) {
	return sql.Open(svc.DBDriver, svc.DBDSN)
}

func (l *Lot) MarkAsTweeted(statusID string) (err error) {
	db, err := openIt()
	if err != nil {
		return
	}
	_, err = db.Exec("UPDATE lots SET tweeted = ? WHERE id = ?", statusID, l.ID)
	return
}

func Next() (l *Lot, err error) {
	db, err := openIt()
	if err != nil {
		return
	}
	var lat, lon sql.NullFloat64
	row := db.QueryRow("SELECT id, lat, lon, address, city, state, zip, tweeted FROM lots WHERE tweeted = 0 ORDER BY id LIMIT 1;")
	if row == nil {
		return nil, fmt.Errorf("No results")
	}
	l = &Lot{}
	err = row.Scan(&l.ID, &lat, &lon, &l.Address, &l.City, &l.State, &l.Zip, &l.Tweeted)
	if lat.Valid && lon.Valid {
		l.Lat = lat.Float64
		l.Lon = lon.Float64
	}

	return
}

func ID(id string) (l *Lot, err error) {
	db, err := openIt()
	if err != nil {
		return
	}
	row := db.QueryRow(`SELECT id, lat, lon, address, city, state, zip, tweeted 
	FROM lots 
	WHERE tweeted = 0 
	AND id = ?
	ORDER BY id 
	LIMIT 1;`, id)
	if row == nil {
		return nil, fmt.Errorf("No results")
	}
	l = &Lot{}
	err = row.Scan(&l.ID, &l.Lat, &l.Lon, &l.Address, &l.City, &l.State, &l.Zip, &l.Tweeted)
	return
}
