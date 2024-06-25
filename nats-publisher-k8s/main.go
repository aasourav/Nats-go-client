package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		// log.Fatal(err)
		log.Println("ERROR in Connection: ", err.Error())
	}
	defer nc.Close()

	// Create a subject (channel)
	subject := "ingress-nginx-controller-6f6cf945d-lfsrc"
	done := make(chan bool)
	// Subscribe to the subject
	go isIAmdone(nc, done)

	for {
		select {
		case <-done:
			fmt.Println("isIAmdone returned true. Terminating main loop.")
			return
		default:
			time.Sleep(100 * time.Millisecond)
			err = nc.Publish(subject, []byte("I am log: "+time.Now().String()))
			if err != nil {
				log.Fatal(err)
			}
			// time.Sleep(1 * time.Second)
		}
	}
}

func isIAmdone(nc *nats.Conn, done chan bool) {
	defer nc.Close()

	for {
		sub, err := nc.SubscribeSync("ingress-nginx-controller-6f6cf945d-lfsrc/keepAlive")
		if err != nil {
			log.Println("subcribe synce err: " + err.Error())
		}
		// time.Sleep(3 * time.Second)
		_, err = sub.NextMsg(1 * time.Second)
		if err != nil {
			log.Println("Nex msg err: " + err.Error())
			done <- true
		}

	}
}
