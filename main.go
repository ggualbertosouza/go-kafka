package main

import (
	"fmt"
	"time"

	"github.com/ggualbertosouza/go-kafka/pubsub"
)

func main() {
	broker := pubsub.NewBroker()

	ch1 := broker.Subscribe("orders")
	go func() {
		for msg := range ch1 {
			fmt.Println("Consumer 1:", msg.Payload)
		}
	}()

	ch2 := broker.Subscribe("orders")
	go func() {
		for msg := range ch2 {
			fmt.Println("Consumer 2:", msg.Payload)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			msg := pubsub.Message{Payload: fmt.Sprintf("Order %d", i)}
			broker.Publish("orders", msg)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
