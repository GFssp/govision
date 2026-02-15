package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQConsumer(ch *amqp.Channel, queue string) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		channel: ch,
		queue:   queue,
	}
}

func (c *RabbitMQConsumer) Consume(
	ctx context.Context,
) (<-chan amqp.Delivery, error) {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return msgs, nil
}
