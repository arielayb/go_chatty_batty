package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type AmqpChannel interface {
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}

type AmqpConnection interface {
	Channel() (AmqpChannel, error)
	Close() error
}

type AmqpDial func(url string) (AmqpConnection, error)

type AmqpConnectionWrapper struct {
	conn *amqp.Connection
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func RmqConnect(rmqpurl string) (AmqpConnection, error) {
	conn, err := amqp.Dial(rmqpurl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

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
	return AmqpConnectionWrapper{conn}, nil
}

func (w AmqpConnectionWrapper) Channel() (AmqpChannel, error) {
	return w.conn.Channel()
}

func (w AmqpConnectionWrapper) Close() error {
	return w.conn.Close()
}
