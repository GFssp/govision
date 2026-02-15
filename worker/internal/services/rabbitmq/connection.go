package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbittMQConnection(connectionString string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
