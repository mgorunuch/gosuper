package gosuper

import (
	"errors"
)

var SuperQueueNotValidTypeErr = errors.New("super queue not valid type")

type SuperQueueConsumer[Type any] struct {
	consumers []func(Type)
}

func (sqCons *SuperQueueConsumer[Type]) Add(consumers ...func(Type)) {
	sqCons.consumers = append(sqCons.consumers, consumers...)
}

func (sqCons *SuperQueueConsumer[Type]) Consume(value any) (err error) {
	val, ok := value.(Type)
	if !ok {
		return SuperQueueNotValidTypeErr
	}

	for _, consumer := range sqCons.consumers {
		consumer(val)
	}

	return err
}
