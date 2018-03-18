package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsulDomain(t *testing.T) {
	Cfg.ConsulDomain = "mydomain"

	assert.Equal(t, ".service.mydomain", ConsulDomain())
}
