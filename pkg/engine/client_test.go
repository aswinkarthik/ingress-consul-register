package engine

import (
	"testing"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestConsulConfig(t *testing.T) {
	config.Cfg.ConsulHost = "10.0.0.100"
	config.Cfg.ConsulPort = 8500

	consulCfg := consulConfig()

	assert.Equal(t, "10.0.0.100:8500", consulCfg.Address)
}
