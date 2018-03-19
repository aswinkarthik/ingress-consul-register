package engine

import (
	"context"
	"fmt"
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/config"
	"github.com/ericchiang/k8s/apis/core/v1"
)

type controllerService struct {
	service *v1.Service
}

func fetchControllerService() *controllerService {
	serviceName := config.ControllerService()
	namespace := config.ControllerServiceNamespace()

	var service v1.Service

	log.Println(fmt.Sprintf("Fetching Service information for service '%s' from namespace %s", serviceName, namespace))
	if err := client.Get(context.Background(), namespace, serviceName, &service); err != nil {
		log.Println(fmt.Sprintf("Failed to fetch %s from namespace %s", serviceName, namespace))
		log.Fatal(err)
	}

	return &controllerService{&service}
}

func (s *controllerService) getIpAddress() string {
	if ipaddress := s.service.GetSpec().GetLoadBalancerIP(); ipaddress != "" {
		return ipaddress
	}
	return s.service.GetSpec().GetClusterIP()
}

func (s *controllerService) getName() string {
	return *s.service.GetMetadata().Name
}

func watchService(fn func()) {
	serviceName := config.ControllerService()
	namespace := config.ControllerServiceNamespace()

	var service v1.Service

	watcher, err := client.Watch(context.Background(), namespace, &service)
	if err != nil {
		log.Println("Error happened while watching service")
		log.Println(err)
	}

	defer watcher.Close()

	for {
		svc := new(v1.Service)
		eventType, err := watcher.Next(svc)

		if err != nil {
			log.Println("Service Watcher encountered an error")
			log.Fatal(err)
		}
		log.Println(fmt.Sprintf("EventType: %s, ServiceName: %s", eventType, svc.GetMetadata().GetName()))
		if svc.GetMetadata().GetName() == serviceName {
			fn()
		}
	}
}
