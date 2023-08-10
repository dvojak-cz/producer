package messagebroker

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func (rbmb *RabbitMQBroker) Close() {
	rbmb.ch.Close()
	rbmb.conn.Close()
}

func NewRabbitMQBroker(connectionData IConnectionData) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(connectionData.ConnectionString())
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQBroker{conn, ch}, nil
}

func (rbmb *RabbitMQBroker) DeclareQueue(name string) error {
	_, err := rbmb.ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (rbmb *RabbitMQBroker) Publish(queueName string, body []byte) error {
	return rbmb.ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
}