package gosuper

import (
	"errors"
)

var SuperQueueNotValidTypeErr = errors.New("super queue not valid type")

type SuperQueueConsumer interface {
	Consume(value any) error
}

type SuperQueueConsumerImpl[Type any] struct {
	consumers []func(Type)
}

func (sqCons *SuperQueueConsumerImpl[Type]) Add(consumers ...func(Type)) {
	sqCons.consumers = append(sqCons.consumers, consumers...)
}

func (sqCons *SuperQueueConsumerImpl[Type]) Consume(value any) (err error) {
	val, ok := value.(Type)
	if !ok {
		return SuperQueueNotValidTypeErr
	}

	for _, consumer := range sqCons.consumers {
		consumer(val)
	}

	return err
}
