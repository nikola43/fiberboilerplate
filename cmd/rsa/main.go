package main

import (
	"log"

	"github.com/nikola43/fiberboilerplate/pkg/utils"
)

func main() {
	// Generate a private key
	err := utils.GenerateRSAPrivateKey()
	if err != nil {
		log.Fatal("GenerateRSAPrivateKey: %v", err.Error())
	}
}
