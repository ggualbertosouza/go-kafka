package core

type Topic struct {
	subscribers []chan Message
}
