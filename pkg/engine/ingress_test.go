package engine

import (
	"testing"

	v1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
	v1 "github.com/ericchiang/k8s/apis/meta/v1"
	"github.com/stretchr/testify/assert"
)

func TestFilterIngressByClass(t *testing.T) {
	nginxClassAnnotation := make(map[string]string, 1)
	nginxClassAnnotation["kubernetes.io/ingress.class"] = "nginx"
	diffClassAnnotation := make(map[string]string, 1)
	diffClassAnnotation["kubernetes.io/ingress.class"] = "nginx-internal"

	expectedIngress := &v1beta1.Ingress{Metadata: &v1.ObjectMeta{Annotations: nginxClassAnnotation}}
	differentClassIngress := &v1beta1.Ingress{Metadata: &v1.ObjectMeta{Annotations: diffClassAnnotation}}

	input := []*v1beta1.Ingress{expectedIngress, differentClassIngress}
	output := filterIngressByClass(input, "nginx")

	assert.Equal(t, 1, len(output))
	assert.Equal(t, expectedIngress, output[0])
}
