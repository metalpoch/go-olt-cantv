package helper

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/metalpoch/go-olt-cantv/model"
)

func ParseGPON(s string) model.Element {
	s = strings.Replace(s, "GPON ", "", 1)
	parts := strings.Split(s, "/")
	shell, err := strconv.Atoi(parts[0])

	if err != nil {
		log.Printf("error parsing %s: %s\n", s, err.Error())
	}

	card, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Printf("error parsing %s: %s\n", s, err.Error())
	}

	port, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Printf("error parsing %s: %s\n", s, err.Error())
	}

	return model.Element{
		Shell: shell,
		Card:  card,
		Port:  port,
	}
}

func BytesToKbps(prevBytes, currBytes, diffDate int) int {
	bps := math.Abs(float64(8*currBytes)-float64(8*prevBytes)) / float64(diffDate)
	return int(math.Round(bps / 1000))
}
