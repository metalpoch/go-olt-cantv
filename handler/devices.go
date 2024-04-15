package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/metalpoch/go-olt-cantv/database"
	"github.com/metalpoch/go-olt-cantv/pkg/snmp"
)

func GetDevices() {
	db := database.DeviceConnect()
	devices, err := handlerDevice(db).FindAll()

	if err != nil {
		log.Fatal("error searching for devices:", err.Error())
	}

	fmt.Println("Sysname\t\tIP\t\tCommunity")
	for _, device := range devices {
		fmt.Printf("%s\t%s\t%s\n", device.Sysname, device.IP, device.Community)
	}
}

func AddDevices(ip string, community string) {
	device := snmp.Sysname(ip, community)

	db := database.DeviceConnect()
	err := handlerDevice(db).Add(device)
	if err != nil {
		log.Println("error saving device:", err.Error())
		os.Exit(1)
	}

}
