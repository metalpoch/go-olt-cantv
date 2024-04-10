package snmp

import (
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

func Sysname(ip, host, community string) model.Device {
	sysname := model.Device{
		IP:        ip,
		Community: community,
	}

	snmp_command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, sysname_oid)
	out, err := exec.Command("ssh", host, snmp_command).Output()
	if err != nil {
		log.Fatalln("error when running snmp command, check IP and Community")
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

func Measurements(ip, host, community string) model.Snmp {
	ports_map := make(map[int]string)
	in_map := make(map[int]int)
	out_map := make(map[int]int)
	bandwidth_map := make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(4)

	go func(oid string) {
		defer wg.Done()

		snmp_command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, oid)
		out, err := exec.Command("ssh", host, snmp_command).Output()
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}

			ports_map[index] = value

		}
	}(ifname_oid)

	go func(oid string) {
		defer wg.Done()

		snmp_command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, oid)
		out, err := exec.Command("ssh", "taccess@161.196.112.163", snmp_command).Output()
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}

			valueInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			in_map[index] = valueInt

		}
	}(bytes_in_oid)

	go func(oid string) {
		defer wg.Done()

		snmp_command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, oid)
		out, err := exec.Command("ssh", "taccess@161.196.112.163", snmp_command).Output()
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}

			valueInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			out_map[index] = valueInt

		}
	}(bytes_out_oid)

	go func(oid string) {
		defer wg.Done()

		snmp_command := fmt.Sprintf("snmpwalk -v 2c -c %s %s %s", community, ip, oid)
		out, err := exec.Command("ssh", "taccess@161.196.112.163", snmp_command).Output()
		if err != nil {
			log.Fatal(err)
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
				log.Fatal(err)
			}

			valueInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			bandwidth_map[index] = valueInt

		}
	}(bandwidth_oid)

	wg.Wait()

	return model.Snmp{
		IfName:    ports_map,
		ByteIn:    in_map,
		ByteOut:   out_map,
		Bandwidth: bandwidth_map,
	}
}
