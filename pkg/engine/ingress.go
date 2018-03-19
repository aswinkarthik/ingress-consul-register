package engine

import (
	"context"
	"fmt"
	"log"

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

func watchIngress(fn func()) {
	var ingress v1beta1.Ingress
	watcher, err := client.Watch(context.Background(), k8s.AllNamespaces, &ingress)
	if err != nil {
		log.Println("Error happened while watching ingresses")
		log.Println(err)
	}
	defer watcher.Close()

	for {
		ing := new(v1beta1.Ingress)
		eventType, err := watcher.Next(ing)
		if err != nil {
			log.Println("Ingress Watcher encountered an error. Exiting")
			log.Fatal(err)
		}
		log.Println(fmt.Sprintf("EventType: %s, IngressName: %s", eventType, ing.GetMetadata().GetName()))
		fn()
	}
}
