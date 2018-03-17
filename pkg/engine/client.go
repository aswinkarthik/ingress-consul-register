package engine

import (
	"github.com/ericchiang/k8s"
)

var client *k8s.Client

func Initialize() error {
	if c, err := k8s.NewInClusterClient(); err != nil {
		return err
	} else {
		client = c
	}
	return nil
}
