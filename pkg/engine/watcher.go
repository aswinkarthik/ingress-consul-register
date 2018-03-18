package engine

import (
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/aswinkarthik93/ingress-consul-register/pkg/utils"
	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
	"github.com/hashicorp/consul/api"
)

type consulService struct {
	Name      string
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
	c.registerToConsul(consulClient)
}

func retrieveTags(ingresses []*v1beta1.Ingress) []string {
	return convertToHosts(ingresses).filterByDomain(config.ConsulDomain()).getTags(config.ConsulDomain(), config.ConsulControllerService())
}

func (c *consulService) registerToConsul(client *api.Client) error {
	if err := client.Agent().ServiceRegister(c.agentServiceRegistration()); err != nil {
		log.Println("An error occurred when registering to consul")
		log.Println(err)
		return err
	}
	return nil
}

func StartWatching() {
	log.Println("Start watching")
}

func (c *consulService) agentServiceRegistration() *api.AgentServiceRegistration {
	return &api.AgentServiceRegistration{
		ID:      c.Name,
		Name:    c.Name,
		Tags:    c.Tags,
		Address: c.IpAddress,
	}
}
