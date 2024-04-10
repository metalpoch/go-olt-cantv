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
	unix INTEGER NOT NULL
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
	device_id INTEGER NOT NULL,
	shell INTEGER NOT NULL,
	card INTEGER NOT NULL,
	port INTEGER NOT NULL,
	FOREIGN KEY (device_id) REFERENCES devices(id),
	UNIQUE (shell, card, port, device_id)
);`

const createCountsTable string = `
CREATE TABLE IF NOT EXISTS tempcounts (
	element_id INTEGER NOT NULL,
	date_id INTEGER NOT NULL,
	bytes_in INTEGER NOT NULL,
	bytes_out INTEGER NOT NULL,
	bandwidth INTEGER NOT NULL,
	FOREIGN KEY (element_id) REFERENCES elements(id),
	FOREIGN KEY (date_id) REFERENCES dates(id)
);`

const createTrafficTable string = `
CREATE TABLE IF NOT EXISTS traffic (
	id INTEGER NOT NULL PRIMARY KEY,
	element_id INTEGER,
	date_id INTEGER,
	kbps_in INTEGER NOT NULL,
	kbps_out INTEGER NOT NULL,
	bandwidth INTEGER NOT NULL,
  FOREIGN KEY (date_id) REFERENCES dates(id),
  FOREIGN KEY (element_id) REFERENCES elements(id)
);`

func createTables(db *sql.DB) error {
	if _, err := db.Exec(createDatesTable); err != nil {
		return err
	}

	if _, err := db.Exec(createDevicesTable); err != nil {
		return err
	}

	if _, err := db.Exec(createCountsTable); err != nil {
		return err
	}

	if _, err := db.Exec(createElementsTable); err != nil {
		return err
	}

	if _, err := db.Exec(createTrafficTable); err != nil {
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
