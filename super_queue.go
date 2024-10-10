package gosuper

import "errors"

type SuperQueue struct {
	consumers []SuperQueueConsumer[any]
}

func (sq *SuperQueue) AddConsumer(consumer SuperQueueConsumer[any]) {
	sq.consumers = append(sq.consumers, consumer)
}

func (sq *SuperQueue) Push(value any) error {
	for _, consumer := range sq.consumers {
		err := consumer.Consume(value)
		if err != nil && !errors.Is(err, SuperQueueNotValidTypeErr) {
			return err
		}
	}

	return nil
}

func NewSuperQueue() *SuperQueue {
	return &SuperQueue{}
}
