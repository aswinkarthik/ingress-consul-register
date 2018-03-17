package config

import "fmt"

type Config struct {
	IngressClass               string
	ConsulDomain               string
	ControllerService          string
	ControllerServiceNamespace string
}

var Cfg Config

func init() {
	Cfg = Config{}
}

func IngressClass() string {
	return Cfg.IngressClass
}

func ConsulDomain() string {
	return fmt.Sprintf(".%s.%s", "consul", Cfg.ConsulDomain)
}
