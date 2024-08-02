package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"heimdall/internal/config"
	"log"
	"log/slog"
	"time"

	rmq "github.com/rabbitmq/amqp091-go"
)

var RmqConn *rmq.Connection //RmqConn is the connection to the rabbitmq instance
// Connect initiates a connection to the rabbitmq instance
func Connect(config *config.Config) (err error) {
	RmqConn, err = rmq.Dial(config.RMQUrl)
	return err
}

// PublishMessage publishes a message to a queue task
func (rp RMQProducer) PublishMessage(message any) error {

	jsonStr, err := json.Marshal(message)
	if err != nil {
		return err
	}

	ch, err := RmqConn.Channel()
	if err != nil {
		return err
	}

	queue, err := ch.QueueDeclare(rp.Queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	pubMessage := rmq.Publishing{
		ContentType: "application/json",
		Timestamp:   time.Now().UTC(),
		Body:        jsonStr,
	}

	return ch.PublishWithContext(context.Background(), "", queue.Name, false, false, pubMessage)
}

// Consume consumes messages from the task queue
func (rc RMQConsumer) Consume() {

	ch, err := RmqConn.Channel()
	if err != nil {
		rmq.Logger.Printf("[Queue]: failed to create channel: %v", err.Error())
	}

	queue, err := ch.QueueDeclare(rc.Queue, true, false, false, false, nil)
	if err != nil {
		rmq.Logger.Printf("[Queue]: failed to declare channel: %v", err.Error())
	}

	messages, err := ch.Consume(queue.Name, "", rc.AutoAck, false, false, false, nil)
	if err != nil {
		rmq.Logger.Printf("[Queue]: failed to consume queue: %s | error : %v", rc.Queue, err.Error())
	}

	go func() {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = errors.WithStack(t)
				default:
					err = fmt.Errorf("unknown error %v", t)
				}
				log.Println(fmt.Sprintf("[worker."+rc.Queue+"] Panic!! %v", err), slog.LevelError, true)
			}
		}()
		for d := range messages {
			rc.MsgHandler(context.Background(), d)
		}
	}()
}

func Nack(deliver rmq.Delivery) {
	if err := deliver.Nack(false, true); err != nil {
		log.Println("[Nack]: failed to nack delivery message")
	}
}

func Ack(deliver rmq.Delivery) {
	if err := deliver.Ack(false); err != nil {
		log.Println("[Ack]: failed to ack delivery message")
	}
}
