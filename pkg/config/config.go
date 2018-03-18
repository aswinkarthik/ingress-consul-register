package config

import "fmt"

type Config struct {
	IngressClass               string
	ConsulDomain               string
	ControllerService          string
	ControllerServiceNamespace string
	ConsulControllerService    string
}

var Cfg Config

func init() {
	Cfg = Config{}
}

func IngressClass() string {
	return Cfg.IngressClass
}

func ConsulDomain() string {
	return fmt.Sprintf(".%s.%s", "service", Cfg.ConsulDomain)
}

func ControllerService() string {
	return Cfg.ControllerService
}

func ControllerServiceNamespace() string {
	return Cfg.ControllerServiceNamespace
}

func ConsulControllerService() string {
	return Cfg.ConsulControllerService
}
