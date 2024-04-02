package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/model"
)

const createDatesTable string = `
CREATE TABLE IF NOT EXISTS dates (
	id INTEGER NOT NULL PRIMARY KEY,
	date INTEGER NOT NULL
);`

const createDevicesTable string = `
CREATE TABLE IF NOT EXISTS devices (
	id INTEGER NOT NULL PRIMARY KEY,
	ip TEXT UNIQUE NOT NULL,
	community TEXT NOT NULL,
	sysname TEXT NOT NULL
);`

const createElementsTable string = `
CREATE TABLE IF NOT EXISTS elements (
	id INTEGER NOT NULL PRIMARY KEY,
	port TEXT NOT NULL,
	device_id INTEGER NOT NULL,
	FOREIGN KEY (device_id) REFERENCES devices(id),
	UNIQUE (port, device_id)
);`

const createMeasurementsTable string = `
CREATE TABLE IF NOT EXISTS measurements (
	id INTEGER NOT NULL PRIMARY KEY,
	bytes_in INTEGER NOT NULL,
	bytes_out INTEGER NOT NULL,
	bandwidth INTEGER NOT NULL,
	date_id INTEGER,
	elements_id INTEGER,
  FOREIGN KEY (date_id) REFERENCES dates(id),
  FOREIGN KEY (elements_id) REFERENCES elements(id)
);`

func createTables(db *sql.DB) error {
	if _, err := db.Exec(createDatesTable); err != nil {
		return err
	}

	if _, err := db.Exec(createDevicesTable); err != nil {
		return err
	}

	if _, err := db.Exec(createElementsTable); err != nil {
		return err
	}

	if _, err := db.Exec(createMeasurementsTable); err != nil {
		return err
	}
	return nil
}

func Connect() (*sql.DB, error) {
	var cfg model.Config = config.LoadConfiguration()
	db, err := sql.Open("sqlite3", cfg.DatabaseFilename)
	if err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}
