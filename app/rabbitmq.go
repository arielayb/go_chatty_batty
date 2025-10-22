package app

import (
	"fmt"
	"log"
	"time"

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

func (r *RmqMsg) RmqConnect(rmqpurl string) {
	conn, err := r.dialer.Dial(rmqpurl)
	r.failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	if err != nil {
		fmt.Errorf("Error: %v", err.Error())
	}

	ch, err := conn.Channel()
	r.failOnError(err, "Failed to open a channel")
	defer ch.Close()

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
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	r.failOnError(err, "Failed to register a consumer")
	// forever := make(chan bool)
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		time.Sleep(time.Duration(time.Second * 4))
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// <-forever
}
