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
	subject := "ingress-nginx-controller-6f6cf945d-lfsrc"

	// Subscribe to the subject
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		// This callback will be invoked when messages are published to the subject
		if string(msg.Data) != "" {
			log.Printf("Received message: %s", string(msg.Data))
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Keep the connection alive
	select {}
}
