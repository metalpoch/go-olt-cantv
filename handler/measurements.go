package handler

import (
	"database/sql"
	"log"
	"strings"
	"sync"
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
		log.Fatalln("error searching for devices:", err.Error())
	}

	if len(devices) == 0 {
		log.Fatalln("no device data to scan")
	}

	unix_time := time.Now().Unix()
	time_string := time.Unix(unix_time, 0).String()
	date_id, err := handlerDate(db).Add(int(unix_time))
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: recorrer de manera asincrona cada dispositivo
	var wg sync.WaitGroup
	wg.Add(len(devices))
	for _, device := range devices {

		go func() {
			log.Printf("getting data from %s at %s\n", device.Sysname, time_string)
			measurements := snmp.Measurements(device.IP, cfg.ProxyHost, device.Community)

			for idx, ifname := range measurements.IfName {

				if !strings.HasPrefix(ifname, "GPON") {
					continue
				}
				gpon := helper.ParseGPON(ifname)

				element := model.Element{
					Shell:    gpon.Shell,
					Card:     gpon.Card,
					Port:     gpon.Port,
					DeviceID: device.ID,
				}

				// find device in db
				elementID, err := handlerElement(db).FindID(element)

				if err != sql.ErrNoRows && err != nil {
					log.Printf("error when searching %s: %s\n", ifname, err.Error())
				}

				// element not exist
				if err == sql.ErrNoRows {
					elementID, err = handlerElement(db).Save(element)
					if err != nil {
						log.Printf("error when trying to save %s with device_id %d: %s\n", ifname, device.ID, err.Error())
						continue
					}
				}

				countDiff, err := handlerCount(db).Add(model.Count{
					ElementID: elementID,
					DateID:    date_id,
					BytesIn:   measurements.ByteIn[idx],
					BytesOut:  measurements.ByteOut[idx],
					Bandwidth: measurements.Bandwidth[idx],
				})

				if err != sql.ErrNoRows && err != nil {
					log.Println(err.Error())
					continue
				}

				if countDiff.ElementID > 0 {
					firstDate, _ := handlerDate(db).Get(countDiff.PrevDateID)
					lastDate, _ := handlerDate(db).Get(countDiff.CurrDateID)

					if _, err := handlerTraffic(db).Add(countDiff, firstDate, lastDate); err != nil {
						log.Printf("error when trying to save the traffic of %s on day %d: %s\n", ifname, date_id, err.Error())
					}
				}

			}
			wg.Done()
		}()

	}
	wg.Wait()
}
