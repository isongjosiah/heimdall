package logic

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"heimdall/internal/config"
	"heimdall/internal/service/queue"
	"log"
)

func InitWorkers(l *Logic, config *config.Config) {
	err := queue.Connect(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	processCommitPull(l.Repository.handleInitialPull)
	processRepositoryAddition(l.Repository.handleRepositoryAddition)

}

type workerHandler func(ctx context.Context, message amqp091.Delivery)

func processRepositoryAddition(handler workerHandler) {

	(queue.RMQConsumer{
		Queue:      "pull-repo",
		MsgHandler: handler,
	}).Consume()
}

// processCommitMonitoring ...
func processCommitPull(handler workerHandler) {

	(queue.RMQConsumer{
		Queue:      "pull-commit",
		MsgHandler: handler,
	}).Consume()

}
