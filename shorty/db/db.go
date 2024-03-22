package db

import (
	"database/sql"
	"errors"
	"shorty/models"
)

var (
	createDb string = `Create table if not exists links (
	original_url string PRIMARY Key,
	short_id string NOT NULL,
	views integer NOT NULL,
	stamp integer NOT NULL,
    published integer NOT NULL)`

	createIdx string = `Create Index short_id_idx on links(short_id)`

	rec *models.Link
)

func CreateTableAndIndex(dbConn *sql.DB) error {
	_, err := dbConn.Exec(createDb)
	if err != nil {
		return err
	}

	_, err = dbConn.Exec(createIdx)
	if err != nil {
		return err
	}

	return nil
}

func AddRecord(rec *models.Link, dbConn *sql.DB) (sql.Result, error) {
	result, err := dbConn.Exec("insert into links (original_irl, short_id, views, stamp, published) Values (?, ?, ?, ?, ?)",
		rec.OriginalUrl, rec.ShortId, rec.Views, rec.Stamp, rec.Published)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateRecord(views int, short_id string, dbConn *sql.DB) (sql.Result, error) {
	result, err := dbConn.Exec("update links SET views = ? where short_id = ?", views, short_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SearchUrl(q string, dbConn *sql.DB) (string, error) {
	var url string
	row := dbConn.QueryRow("Select original_url from links where short_id = ?", q)
	if err := row.Scan(url); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
	}
	return url, nil
}
