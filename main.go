package main

import (
	"fmt"
	"os"

	"github.com/metalpoch/go-olt-cantv/database"
	"github.com/metalpoch/go-olt-cantv/handler"
)

func main() {
	var args []string
	var action string

	if len(os.Args) == 1 {
		fmt.Println("error: You need to specify an action [device|scan]")
		os.Exit(0)
	}

	args = os.Args[1:]
	action = args[0]

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	switch action {
	case "device":
		if len(args) < 2 {
			fmt.Println("error: you need to specify an option [get|add]")
			return
		}
		switch args[1] {
		case "get":
			handler.GetDevices(db)
		case "add":
			if len(args) < 4 {
				fmt.Println("error: it is necessary to specify IP and community, in that same order")
				return
			}
			var ip string = args[2]
			var community string = args[3]
			handler.AddDevices(db, ip, community)
		default:
			fmt.Println("error: invalid option, you need to specify an option [get|add]")
		}

	case "scan":
		handler.GetMeasurements(db)
	default:
		fmt.Println("error: invalid action")
	}

}
