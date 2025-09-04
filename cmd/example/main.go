package main

import (
	"log"

	"github.com/peer-coding/pipe/pkg/ulid"
)

func main() {
	log.Println("starting app...")

	ulid := ulid.New()
	log.Printf("ulid=%s\n", ulid)

	log.Println("finished app!")
}
