package engine

import (
	"testing"

	"github.com/ericchiang/k8s/apis/core/v1"
	meta "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/stretchr/testify/assert"
)

func TestGetIpAddressReturnsClusterIp(t *testing.T) {
	ipaddress := "10.2.2.30"
	s := &controllerService{
		service: &v1.Service{Spec: &v1.ServiceSpec{ClusterIP: &ipaddress}},
	}
	assert.Equal(t, ipaddress, s.getIpAddress())
}

func TestGetIpAddressReturnsLoadBalancerIpIfPresent(t *testing.T) {
	ipaddress := "10.2.2.30"
	loadbalancerIP := "10.200.10.5"
	s := &controllerService{
		service: &v1.Service{Spec: &v1.ServiceSpec{
			ClusterIP:      &ipaddress,
			LoadBalancerIP: &loadbalancerIP,
		}},
	}
	assert.Equal(t, loadbalancerIP, s.getIpAddress())
}

func TestGetName(t *testing.T) {
	serviceName := "router"
	metadata := &meta.ObjectMeta{Name: &serviceName}
	s := &controllerService{
		service: &v1.Service{Metadata: metadata},
	}

	assert.Equal(t, "router", s.getName())
}
