package engine

import (
	"testing"

	"github.com/ericchiang/k8s/apis/extensions/v1beta1"
	v1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/stretchr/testify/assert"
)

func TestConvertToHosts(t *testing.T) {
	annotation := make(map[string]string, 1)
	annotation["kubernetes.io/ingress.class"] = "nginx"
	host1 := "abc.domain.com"
	host2 := "def.domain.com"
	host3 := "ijk.domain.com"

	ing1 := &v1beta1.Ingress{
		Metadata: &v1.ObjectMeta{Annotations: annotation},
		Spec:     &v1beta1.IngressSpec{Rules: []*v1beta1.IngressRule{&v1beta1.IngressRule{Host: &host1}, &v1beta1.IngressRule{Host: &host2}}},
	}
	ing2 := &v1beta1.Ingress{
		Metadata: &v1.ObjectMeta{Annotations: annotation},
		Spec:     &v1beta1.IngressSpec{Rules: []*v1beta1.IngressRule{&v1beta1.IngressRule{Host: &host3}}},
	}

	input := []*v1beta1.Ingress{ing1, ing2}
	output := convertToHosts(input)
	expectedHosts := hosts([]*string{&host1, &host2, &host3})

	assert.Equal(t, 3, len(output))
	assert.Equal(t, expectedHosts, output)
}

func TestFilterByDomain(t *testing.T) {
	host1 := "valid.internal.domain"
	host2 := "invalid_sub.internal.domain"
	host3 := "valid.internal.randomdomain"

	input := hosts([]*string{&host1, &host2, &host3})

	actual := input.filterByDomain("internal.domain")
	expected := hosts([]*string{&host1})

	assert.Equal(t, expected, actual)
}

func TestHostsSize(t *testing.T) {
	host1 := "host1.internal.domain"
	host2 := "host2.internal.domain"
	host3 := "host3.internal.domain"

	input := hosts([]*string{&host1, &host2, &host3})

	assert.Equal(t, 3, input.size())
}

func TestHostsItems(t *testing.T) {
	host1 := "host1.internal.domain"
	host2 := "host2.internal.domain"
	host3 := "host3.internal.domain"

	hostArray := []*string{&host1, &host2, &host3}

	input := hosts(hostArray)

	assert.Equal(t, hostArray, input.items())
}

func TestGetTags(t *testing.T) {
	host1 := "host1.router.internal.domain"
	host2 := "host2.router.internal.domain"
	host3 := "host3.router.internal.domain"

	input := hosts([]*string{&host1, &host2, &host3})

	actual := input.getTags("internal.domain", "router")
	expected := []string{"host1", "host2", "host3"}

	assert.Equal(t, expected, actual)

	actual = input.getTags(".internal.domain", "router")
	assert.Equal(t, expected, actual)
}
