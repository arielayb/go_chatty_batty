package app

import (
	//"fmt"
	//"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AmqpChannel interface {
	ExchangeDeclare(
		name, kind string,
		durable, autoDelete, internal, noWait bool,
		args amqp.Table,
	) error
	QueueDeclare(
		name string,
		durable, autoDelete, exclusive, noWait bool,
		args amqp.Table,
	) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	Consume(
		queue, consumer string,
		autoAck, exclusive, noLocal, noWait bool,
		args amqp.Table,
	) (<-chan amqp.Delivery, error)
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

type AmqpConnection interface {
	Channel() (*amqp.Connection, error)
	Close() error
}

type AmqpDialer interface {
	Dial(url string) (*amqp.Connection, error)
}

type RealRmqDialer struct{}

func (r *RealRmqDialer) Dial(url string) (*amqp.Connection, error) {
	return amqp.Dial(url)
}
