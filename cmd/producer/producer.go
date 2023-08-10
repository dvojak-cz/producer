package main

import (
	conf "github.com/dvojak-cz/producer/pkg/conf"
	mb "github.com/dvojak-cz/producer/pkg/messageBroker"

	log "github.com/sirupsen/logrus"
	logging "github.com/dvojak-cz/producer/pkg/log"
)


var (
	broker *mb.RabbitMQBroker
)

func main() {
	log.Infoln("Starting producer")
	broker, err := mb.NewRabbitMQBroker(
		mb.NewConnectionDataRMQ(
			conf.MessageBrokerHost,
			conf.MessageBrokerPort,
			conf.MessageBrokerUser,
			conf.MessageBrokerPassword,
		),
	)
	logging.PanicOnError(err, "Error creating broker")
	defer broker.Close()
	log.Infoln("Broker created")

	err = broker.DeclareQueue(conf.MessageBrokerQueue)
	logging.PanicOnError(err, "Error declaring queue")
	log.Infoln("Queue declared")

	err = broker.Publish(conf.MessageBrokerQueue, []byte("Hello World!"))
}
