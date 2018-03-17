package engine

import (
	"log"

	"github.com/aswinkarthik93/ingress-consul-register/pkg/utils"
)

func RunOnce() {
	log.Println("Initiating first run")
	ingresses := fetchIngresses()
	utils.PrettyPrint(ingresses)
}

func StartWatching() {
	log.Println("Start watching")
}
