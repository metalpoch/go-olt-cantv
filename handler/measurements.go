package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/model"
	helper "github.com/metalpoch/go-olt-cantv/pkg"
	"github.com/metalpoch/go-olt-cantv/pkg/snmp"
)

func GetMeasurements(db *sql.DB) {
	var cfg model.Config = config.LoadConfiguration()
	devices, err := handlerDevice(db).FindAll()
	if err != nil {
		fmt.Println("error searching for devices:", err.Error())
		os.Exit(1)
	}
	if len(devices) == 0 {
		fmt.Println("no device data to scan")
		os.Exit(1)
	}

	var unix_time uint = uint(time.Now().Unix())
	date_id, err := handlerDate(db).Add(model.Date{Date: unix_time})
	if err != nil {
		fmt.Printf("error search save the date %d: %s\n", unix_time, err.Error())
		os.Exit(1)
	}

	for _, device := range devices {
		log.Printf("getting data from %s at %s\n", device.Sysname, time.Unix(int64(unix_time), 0).String())
		measurements := snmp.Measurements(device.IP, cfg.ProxyHost, device.Community)
		for idx, ifname := range measurements.IfName {
			var element_id uint
			var gpon model.Element
			newElement := model.Element{}

			if !strings.HasPrefix(ifname, "GPON") {
				continue
			}
			gpon = helper.ParseGPON(ifname)

			element := model.Element{
				Shell:    gpon.Shell,
				Card:     gpon.Card,
				Port:     gpon.Port,
				DeviceID: device.ID,
			}

			// find device in db
			elementFound, err := handlerElement(db).Find(element)

			if err != sql.ErrNoRows && err != nil {
				log.Printf("error when searching %s: %s\n", ifname, err.Error())
			}

			// element not exist
			if err == sql.ErrNoRows {
				newElement, err = handlerElement(db).Save(element)
				if err != nil {
					log.Printf("error when trying to save %s with device_id %d: %s\n", ifname, device.ID, err.Error())
				}
				element_id = newElement.ID
			} else {
				element_id = elementFound.ID
			}

			measurement := model.SaveMeasurement{
				ByteIn:    measurements.ByteIn[idx],
				ByteOut:   measurements.ByteOut[idx],
				Bandwidth: uint16(measurements.Bandwidth[idx]),
				DateID:    uint(date_id),
				ElementID: element_id,
			}

			// save measurement
			err = handlerMeasurement(db).Save(measurement)
			if err != nil {
				log.Printf("error when trying to save the measurement of %s on day %d: %s\n", ifname, date_id, err.Error())
			}
		}
	}
}
