# Go Pub/Sub (kafka-like)

# Overview
This project is a from-scratch implementation of Pub/Sub sustem in Go, inspired by system like 
kafka, but focused in learning concurrency, goroutines and channels.

The goal is NOT to recreate kafka, but to deeply understand:
- Concurrency in Go
- Communication via Channels
- Goroutine lifecycle
- Backpressure and buffering
- Fan-out message distribution

---
# Core concepts

## Broker

The **Broker** is the central component of the system.

It is responsible for:
- Managing topics
- Receiving messages from topics
- Distributing messages to subscribers (consumers)

## Producer

A **Producer** is reponsible for:
- Sending messages for specific topic

It does know:
- Who will consume the message
- How many consumers exist

This creating decoupling between producers and consumers.

## Consumers (Subscribers)

A **Consumer**:
- Subscribe to a topic
- Receive messages asynchronously

Each consumer processes messages independently using it own goroutine.

## Topic

A **Topic** is a logical channel of communication.
- Producer publish message to a topic
- Consumers subscribe to a topic

In this project:
- Each topic will have multiple subscribers
- Messages will be broadcasted (fan-out)

---
# Message flow

Producer -> Broker -> Topic -> Consumer

1. Producer send message to a topic
2. Broker receive the message
3. Broker distributes the message to all subscribers of that topic
4. Each consumer processes the message independently

---
# Roadmap

## V1 - Basic pub/sub
- [x] Create broker structure
- [x] Implement producer
- [x] Implement Subscribe
- [x] Basic fan-out (1 message -> n consumer)
- [x] Use channels for communication

## V2 - Concurrency Safety
- [ ] Protect shared state (topic map)
- [ ] Introduce `sync.Mutex`
- [ ] Prevent race condition

## V3 - Buffered channels
- [ ] Buffered channels
- [ ] Define behavior for slow consumers:
    - block produces?
    - drop messages?
- [ ] Explore system limits

## V4 - Consumer lifecycle
- [ ] Unsubscribe mechanism
- [ ] Channel closing strategy
- [ ] Prevent goroutine leaks
