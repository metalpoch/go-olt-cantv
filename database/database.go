package database

import (
	"database/sql"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/model"
)

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
	shell INTEGER NOT NULL,
	card INTEGER NOT NULL,
	port INTEGER NOT NULL,
	UNIQUE (shell, card, port)
);`

const createCountsTable string = `
CREATE TABLE IF NOT EXISTS tempcounts (
	element_id INTEGER NOT NULL,
	date INTEGER NOT NULL,
	bytes_in INTEGER NOT NULL,
	bytes_out INTEGER NOT NULL,
	bandwidth INTEGER NOT NULL,
	FOREIGN KEY (element_id) REFERENCES elements(id)
);`

const createTrafficTable string = `
CREATE TABLE IF NOT EXISTS traffic (
	id INTEGER NOT NULL PRIMARY KEY,
	element_id INTEGER,
	date INTEGER,
	kbps_in INTEGER NOT NULL,
	kbps_out INTEGER NOT NULL,
	bandwidth INTEGER NOT NULL,
  FOREIGN KEY (element_id) REFERENCES elements(id)
);`

func createMeasurementsTables(db *sql.DB) error {
	if _, err := db.Exec(createElementsTable); err != nil {
		return err
	}

	if _, err := db.Exec(createCountsTable); err != nil {
		return err
	}

	if _, err := db.Exec(createTrafficTable); err != nil {
		return err
	}

	return nil
}

func connect(db_name string) *sql.DB {
	var cfg model.Config = config.LoadConfiguration()
	db, err := sql.Open("sqlite3", path.Join(cfg.DirDB, db_name))
	if err != nil {
		panic(err)
	}
	return db
}

func MeasurementConnect(sysname string) *sql.DB {
	db := connect(sysname)
	if err := createMeasurementsTables(db); err != nil {
		panic(err)
	}
	return db
}

func DeviceConnect() *sql.DB {
	db := connect("devices")
	if _, err := db.Exec(createDevicesTable); err != nil {
		panic(err)
	}
	return db
}
