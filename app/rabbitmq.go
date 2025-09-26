package app

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

/*type AmqpChannel interface {
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
}*/

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

func (r *RmqMsg) RmqConnect(rmqpurl string) error {
	conn, err := r.dialer.Dial(rmqpurl)
	r.failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	if err != nil {
		fmt.Errorf("Error: %v", err.Error())
	}

	/*ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	//	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")*/
	r.RmqConn = conn

	return nil
}
