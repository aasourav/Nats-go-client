package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create a subject (channel)
	subject := "ingress-nginx-controller-6f6cf945d-lfsrc/keepAlive"

	// Subscribe to the subject
	err = nc.Publish(subject, []byte("Hello, NATS!"))

	if err != nil {
		log.Fatal(err)
	}

	// Keep the connection alive
	select {}
}
