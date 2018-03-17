package engine

import (
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/aswinkarthik93/ingress-consul-register/pkg/utils"
)

func RunOnce() {
	log.Println("Initiating first run")
	ingresses := fetchIngresses()
	utils.PrettyPrint(ingresses)

	tags := convertToHosts(ingresses).filterByDomain(config.ConsulDomain()).getTags(config.ConsulDomain(), "")
	utils.PrettyPrint(tags)
}

func StartWatching() {
	log.Println("Start watching")
}
