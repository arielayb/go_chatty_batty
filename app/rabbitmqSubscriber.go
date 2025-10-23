package app

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqSub struct {
	AmqpConn *amqp.Connection
}

func (r *RmqSub) failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (r *RmqSub) RmqSubscriber() {
	ch, err := r.AmqpConn.Channel()
	r.failOnError(err, "Failed to open a channel")
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		"test", // name
		true,   // durable
		true,   // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	r.failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	r.failOnError(err, "Failed to register a consumer")
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		time.Sleep(5 * time.Duration(time.Minute))
	}
}
