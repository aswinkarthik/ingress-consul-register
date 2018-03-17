package engine

import (
	"context"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/ericchiang/k8s"
	v1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

const ingressClassKey = "kubernetes.io/ingress.class"

func fetchIngresses() []*v1beta1.Ingress {
	var ingresses v1beta1.IngressList
	client.List(context.Background(), k8s.AllNamespaces, &ingresses)
	return filterIngressByClass(ingresses.GetItems(), config.IngressClass())
}

func filterIngressByClass(ingresses []*v1beta1.Ingress, class string) []*v1beta1.Ingress {
	result := make([]*v1beta1.Ingress, len(ingresses))
	counter := 0

	for _, ing := range ingresses {
		if ing.GetMetadata().GetAnnotations()[ingressClassKey] == class {
			result[counter] = ing
			counter++
		}
	}

	return result[:counter]
}
