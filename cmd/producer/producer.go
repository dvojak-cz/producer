package main

import (
	env "producer/internal/env"
	mb "producer/pkg/messageBroker"

	log "github.com/sirupsen/logrus"
	logging "producer/internal/log"
)


var (
	broker *mb.RabbitMQBroker
)

func main() {
	log.Infoln("Starting producer")
	broker, err := mb.NewRabbitMQBroker(
		mb.NewConnectionDataRMQ(
			env.MessageBrokerHost,
			env.MessageBrokerPort,
			env.MessageBrokerUser,
			env.MessageBrokerPassword,
		),
	)
	logging.PanicOnError(err, "Error creating broker")
	defer broker.Close()
	log.Infoln("Broker created")

	err = broker.DeclareQueue(env.MessageBrokerQueue)
	logging.PanicOnError(err, "Error declaring queue")
	log.Infoln("Queue declared")

	err = broker.Publish(env.MessageBrokerQueue, []byte("Hello World!"))
}
