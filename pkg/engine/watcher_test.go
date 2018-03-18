package engine

import (
	"testing"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
	v1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveTags(t *testing.T) {
	config.Cfg.ConsulDomain = "uat"
	config.Cfg.ConsulControllerService = "internal"

	annotation := make(map[string]string, 1)
	annotation["kubernetes.io/ingress.class"] = "nginx-internal"
	host1 := "payment.internal.service.uat"
	host2 := "dashboard.external.nginx" //Not be registered in consul
	host3 := "ticket.internal.service.uat"

	ing1 := &v1beta1.Ingress{
		Metadata: &v1.ObjectMeta{Annotations: annotation},
		Spec:     &v1beta1.IngressSpec{Rules: []*v1beta1.IngressRule{&v1beta1.IngressRule{Host: &host1}, &v1beta1.IngressRule{Host: &host2}}},
	}
	ing2 := &v1beta1.Ingress{
		Metadata: &v1.ObjectMeta{Annotations: annotation},
		Spec:     &v1beta1.IngressSpec{Rules: []*v1beta1.IngressRule{&v1beta1.IngressRule{Host: &host3}}},
	}

	input := []*v1beta1.Ingress{ing1, ing2}

	actualTags := retrieveTags(input)
	expectedTags := []string{"payment", "ticket"}

	assert.Equal(t, expectedTags, actualTags)
}
