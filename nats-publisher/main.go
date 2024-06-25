package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println(err.Error())
	}
	defer nc.Close()

	// Create a subject (channel)
	subject := "ingress-nginx-controller-6f6cf945d-lfsrc/keepAlive"

	// Subscribe to the subject
	for {
		err = nc.Publish(subject, []byte("send me data"))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// Keep the connection alive
	// select {}
}
