package ssh

import (
	"log"

	"github.com/melbahja/goph"
	"github.com/metalpoch/go-olt-cantv/model"
)

func ClientSSH(config model.Config) *goph.Client {
	auth, err := goph.Key(config.SSHPrivateKey, config.SSHPrivatePassw)
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New(config.ProxyUser, config.ProxyHost, auth)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
