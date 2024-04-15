package main

import (
	"fmt"
	"os"

	"github.com/metalpoch/go-olt-cantv/config"
	"github.com/metalpoch/go-olt-cantv/handler"
	"github.com/metalpoch/go-olt-cantv/model"
	helper "github.com/metalpoch/go-olt-cantv/pkg"
)

func main() {
	var cfg model.Config = config.LoadConfiguration()
	var args []string
	var action string

	if len(os.Args) == 1 {
		fmt.Println("error: You need to specify an action [device|scan]")
		os.Exit(0)
	}

	args = os.Args[1:]
	action = args[0]

	helper.Mkdir(cfg.DirDB)

	switch action {
	case "device":
		if len(args) < 2 {
			fmt.Println("error: you need to specify an option [get|add]")
			return
		}
		switch args[1] {
		case "get":
			handler.GetDevices()
		case "add":
			if len(args) < 4 {
				fmt.Println("error: it is necessary to specify IP and community, in that same order")
				return
			}
			ip := args[2]
			community := args[3]
			handler.AddDevices(ip, community)
		default:
			fmt.Println("error: invalid option, you need to specify an option [get|add]")
		}

	case "scan":
		handler.GetMeasurements()
	default:
		fmt.Println("error: invalid action")
	}

}
