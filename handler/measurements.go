package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/model"
	snmp "github.com/metalpoch/go-olt-cantv/pkg"
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
		for idx, port := range measurements.IfName {
			var element_id uint
			newElement := model.Elements{}
			element := model.Elements{
				Port:     port,
				DeviceID: device.ID,
			}

			// find device in db
			elementFound, err := handlerElement(db).Find(element)

			if err != sql.ErrNoRows && err != nil {
				fmt.Printf("error when searching %s: %s\n", port, err.Error())
			}

			// element not exist
			if err == sql.ErrNoRows {
				newElement, err = handlerElement(db).Save(element)
				if err != nil {
					fmt.Printf("error when trying to save %s with device_id %d: %s\n", port, device.ID, err.Error())
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
				fmt.Printf("error when trying to save the measurement of %d on day %d: %s\n", element_id, date_id, err.Error())
			}
		}
	}
}
