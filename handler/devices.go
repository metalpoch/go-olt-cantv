package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/model"
	"github.com/metalpoch/go-olt-cantv/pkg/snmp"
)

func GetDevices(db *sql.DB) {
	devices, err := handlerDevice(db).FindAll()

	if err != nil {
		log.Fatal("error searching for devices:", err.Error())
	}

	fmt.Println("ID\t\tIP\tCommunity\tSysname")
	for _, device := range devices {
		fmt.Printf("%d\t%s\t%s\t%s\n", device.ID, device.IP, device.Community, device.Sysname)
	}
}

func AddDevices(db *sql.DB, ip string, community string) {
	var cfg model.Config = config.LoadConfiguration()
	device := snmp.Sysname(ip, cfg.ProxyHost, community)
	err := handlerDevice(db).Add(device)
	if err != nil {
		log.Println("error saving device:", err.Error())
		os.Exit(1)
	}

}
