package app

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqMsg struct {
	dialer  AmqpDialer
	RmqConn *amqp.Connection
}

func NewRmqClient(dialer AmqpDialer) *RmqMsg {
	return &RmqMsg{dialer: dialer}
}

func (r *RmqMsg) failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (r *RmqMsg) RmqConnect(rmqpurl string) *amqp.Connection {
	conn, err := r.dialer.Dial(rmqpurl)
	r.failOnError(err, "Failed to connect to RabbitMQ")

	_, err = conn.Channel()
	r.failOnError(err, "Failed to open a channel")

	log.Println("Successfully connected to RabbitMQ.")

	return conn
}
