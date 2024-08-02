package queue

import (
	"context"
	rmq "github.com/rabbitmq/amqp091-go"
)

// RMQProducer represents an instance that is able to publish a message
// to a rabbitmq queue
type RMQProducer struct {
	Queue           string
	DeadLetterQueue string
}

// RMQConsumer represents an instance that is able to consume a message
// from a rabbitqm queue
type RMQConsumer struct {
	Queue      string
	MsgHandler func(ctx context.Context, msg rmq.Delivery)
	AutoAck    bool
}
