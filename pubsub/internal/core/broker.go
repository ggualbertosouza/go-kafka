package core

type Broker struct {
	topics map[string]*Topic
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string]*Topic),
	}
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

func (b *Broker) Subscribe(topicName string) <-chan Message {
	topic := b.getOrCreateTopic(topicName)

	ch := make(chan Message)
	topic.subscribers = append(topic.subscribers, ch)

	return ch
}

func (b *Broker) getOrCreateTopic(topicName string) *Topic {
	topic, exists := b.topics[topicName]
	if !exists {
		topic = &Topic{}
		b.topics[topicName] = topic
	}

	return topic
}
