package main

import (
	"fmt"
	"time"
)

type Message struct {
	Payload any
}

type Topic struct {
	subscribers []chan Message
}

type Broker struct {
	topics map[string]*Topic
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]*Topic),
	}
}

func (b *Broker) Subscribe(topicName string) <-chan Message {
	topic, exists := b.topics[topicName]
	if !exists {
		topic = &Topic{}
		b.topics[topicName] = topic
	}

	ch := make(chan Message)
	topic.subscribers = append(topic.subscribers, ch)

	return ch
}

func (b *Broker) Publish(topicName string, msg Message) {
	topic, exists := b.topics[topicName]
	if !exists {
		return
	}

	for _, sub := range topic.subscribers {
		sub <- msg
	}
}

func main() {
	broker := NewBroker()

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
			msg := Message{Payload: fmt.Sprintf("Order %d", i)}
			broker.Publish("orders", msg)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
