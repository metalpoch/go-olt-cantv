package handler

import (
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/metalpoch/go-olt-cantv/database"
	"github.com/metalpoch/go-olt-cantv/model"
	helper "github.com/metalpoch/go-olt-cantv/pkg"
	"github.com/metalpoch/go-olt-cantv/pkg/snmp"
)

func GetMeasurements() {
	db_devices := database.DeviceConnect()
	devices, err := handlerDevice(db_devices).FindAll()

	if err != nil {
		log.Fatalln("error searching for devices:", err.Error())
	}

	if len(devices) == 0 {
		log.Fatalln("no device data to scan")
	}

	var wg sync.WaitGroup
	wg.Add(len(devices))

	for _, device := range devices {
		go func() {
			defer wg.Done()

			db_sysname := database.MeasurementConnect(device.Sysname)
			if err != nil {
				log.Fatalln(err)
			}

			unix_time := time.Now().Unix()
			measurements := snmp.Measurements(device)

			for idx, ifname := range measurements.IfName {

				if !strings.HasPrefix(ifname, "GPON") {
					continue
				}

				gpon := helper.ParseGPON(ifname)

				element := model.Element{
					Shell: gpon.Shell,
					Card:  gpon.Card,
					Port:  gpon.Port,
				}

				// find device in db
				elementID, err := handlerElement(db_sysname).FindID(element)

				if err != sql.ErrNoRows && err != nil {
					log.Printf("error when searching %s: %s\n", ifname, err.Error())
				}

				// element not exist
				if err == sql.ErrNoRows {
					elementID, err = handlerElement(db_sysname).Save(element)
					if err != nil {
						log.Printf("error when trying to save %s with device_id %d: %s\n", ifname, device.ID, err.Error())
						continue
					}
				}

				countDiff, err := handlerCount(db_sysname).Add(model.Count{
					ElementID: elementID,
					Date:      int(unix_time),
					BytesIn:   measurements.ByteIn[idx],
					BytesOut:  measurements.ByteOut[idx],
					Bandwidth: measurements.Bandwidth[idx],
				})

				if err != sql.ErrNoRows && err != nil {
					log.Println(err.Error())
					continue
				}

				if countDiff.ElementID > 0 {
					if _, err := handlerTraffic(db_sysname).Add(countDiff); err != nil {
						log.Printf("error when trying to save the traffic of %s on day %d: %s\n", ifname, countDiff.CurrDate, err.Error())
					}
				}

			}

		}()

	}
	wg.Wait()
}
