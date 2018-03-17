package utils

import (
	"encoding/json"
	"log"
)

func PrettyPrint(v interface{}) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Could not pretty print given Interface")
		return
	}

	log.Println(string(data))
}
