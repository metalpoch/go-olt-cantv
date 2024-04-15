package snmp

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"

	"github.com/metalpoch/go-olt-cantv/model"
)

const (
	sysname_oid   = "1.3.6.1.2.1.1.5.0"
	ifname_oid    = "1.3.6.1.2.1.31.1.1.1.1"
	bytes_in_oid  = "1.3.6.1.4.1.2011.6.128.1.1.4.21.1.15"
	bytes_out_oid = "1.3.6.1.4.1.2011.6.128.1.1.4.21.1.30"
	bandwidth_oid = "1.3.6.1.2.1.31.1.1.1.15"
)

func shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func Sysname(ip, community string) model.Device {
	sysname := model.Device{
		IP:        ip,
		Community: community,
	}

	command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, sysname_oid)
	out, stderr, err := shellout(command)

	if err != nil || stderr != "" {
		log.Fatalf("device: %s - %s | stderr: %s | err: %w\n", ip, community, stderr, err)
	}

	rows := strings.Split(string(out), "\n")
	for _, row := range rows {
		if len(row) < 1 {
			break
		}
		sysname.Sysname = strings.Split(row, "STRING: ")[1]
	}

	return sysname
}

func Measurements(device model.Device) model.Snmp {
	oids := [4]string{ifname_oid, bytes_in_oid, bytes_out_oid, bandwidth_oid}
	ports_map := make(map[int]string)
	in_map := make(map[int]int)
	out_map := make(map[int]int)
	bandwidth_map := make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(4)

	for _, oid := range oids {

		go func(oid string) {
			defer wg.Done()
			command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", device.Community, device.IP, oid)
			out, stderr, err := shellout(command)

			if err != nil || stderr != "" {
				log.Printf("device: %s | stderr: %s | err: %s\n", device.Sysname, stderr, err.Error())
			}

			rows := strings.Split(string(out), "\n")
			for _, row := range rows {
				if len(row) < 1 {
					break
				}

				var value string
				splited := strings.Split(row, " = ")
				index_part := strings.Split(splited[0], ".")
				value = strings.Split(splited[1], ": ")[1]

				index, err := strconv.Atoi(index_part[len(index_part)-1])
				if err != nil {
					log.Println(device.Sysname, "-", err, string(out))
				}

				if oid == ifname_oid {
					ports_map[index] = value
				} else {
					valueInt, err := strconv.Atoi(value)
					if err != nil {
						log.Println(device.Sysname, "-", err, string(out))
					}

					switch oid {
					case bytes_in_oid:
						in_map[index] = valueInt
					case bytes_out_oid:
						out_map[index] = valueInt
					case bandwidth_oid:
						bandwidth_map[index] = valueInt
					}
				}
			}
		}(oid)
	}

	wg.Wait()

	return model.Snmp{
		IfName:    ports_map,
		ByteIn:    in_map,
		ByteOut:   out_map,
		Bandwidth: bandwidth_map,
	}
}
