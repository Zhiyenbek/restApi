package main

import (
	"log"

	"github.com/Zhiyenbek/restApi"
)

func main() {
	srv := new(restApi.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
