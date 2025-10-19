package app

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqPubMsg struct {
	dialer  AmqpDialer
	RmqConn *amqp.Connection
}

func NewRmqPublisher(dialer AmqpDialer) *RmqPubMsg {
	return &RmqPubMsg{dialer: dialer}
}

func (r *RmqPubMsg) failOnPublishError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (r *RmqPubMsg) RmqPublish(rmqpurl string) {
	// const timedMessages = 43200
	// var processedMsgs int

	conn, err := r.dialer.Dial(rmqpurl)
	r.failOnPublishError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	r.failOnPublishError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	r.failOnPublishError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		for {
			body := "escape!"
			err = ch.PublishWithContext(ctx,
				"logs", // exchange
				"",     // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			r.failOnPublishError(err, "Failed to publish a message")

			log.Printf(" [x] Sent %s", body)
		}
	}()
}
