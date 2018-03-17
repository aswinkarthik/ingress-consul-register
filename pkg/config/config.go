package config

type Config struct {
	IngressClass string
}

var Cfg Config

func init() {
	Cfg = Config{}
}

func IngressClass() string {
	return Cfg.IngressClass
}
