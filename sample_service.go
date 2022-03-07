package main

import (
	"github.com/n1zea144/sampleservice/graphdb"
	"log"
	"os"
)

func run(gc graphdb.RequestRepository) error {
	records, err := gc.GetRequests()
	if err != nil {
		return err
	}
	log.Println("Requests: ", records)
	return nil
}

func main() {
	gc, err := graphdb.NewRequestRepository(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatalln("Error calling graphdb.NewRequestClient(), aborting:", err)
	}

	if err := run(gc); err != nil {
		os.Exit(1)
	}
	log.Println("Exiting sample-service...")
}
