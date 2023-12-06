package main

import (
	"c/dev/projects/http-rest-api/internal/app/apiserver"
	"log"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
