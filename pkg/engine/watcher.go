package engine

import (
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/aswinkarthik93/ingress-consul-register/pkg/utils"
	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

type consulService struct {
	Tags      []string
	IpAddress string
}

func RunOnce() {
	log.Println("Initiating first run")
	ingresses := fetchIngresses()
	tags := retrieveTags(ingresses)
	service := fetchControllerService()
	c := &consulService{Tags: tags, IpAddress: service.getIpAddress()}
	log.Println("Following service will be registered in consul")
	utils.PrettyPrint(c)
}

func retrieveTags(ingresses []*v1beta1.Ingress) []string {
	return convertToHosts(ingresses).filterByDomain(config.ConsulDomain()).getTags(config.ConsulDomain(), config.ConsulControllerService())
}

func StartWatching() {
	log.Println("Start watching")
}
