package engine

import (
	"fmt"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/ericchiang/k8s"
	"github.com/hashicorp/consul/api"
)

var client *k8s.Client
var consulClient *api.Client

func Initialize() error {
	if c, err := k8s.NewInClusterClient(); err != nil {
		return err
	} else {
		client = c
	}
	if c, err := api.NewClient(consulConfig()); err != nil {
		return err
	} else {
		consulClient = c
	}
	return nil
}

func consulConfig() *api.Config {
	conf := api.DefaultConfig()
	conf.Address = fmt.Sprintf("%s:%d", config.ConsulHost(), config.ConsulPort())
	return conf
}
