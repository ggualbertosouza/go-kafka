package pubsub

import "github.com/ggualbertosouza/go-kafka/pubsub/internal/core"

type Message = core.Message

func NewBroker() *core.Broker {
	return core.NewBroker()
}
