package app

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RmqPubMsg struct {
	AmqpConn *amqp.Connection
}

func (r *RmqPubMsg) failOnPublishError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (r *RmqPubMsg) RmqPublish() {
	// simulate a 12 hour message stream
	const timedMessages = 172800
	var processedMsgs int
	const messageRate = 5

	ch, err := r.AmqpConn.Channel()
	r.failOnPublishError(err, "Failed to open a channel")
	//defer ch.Close()

	q, err := ch.QueueDeclare(
		"test", // name
		true,   // durable
		true,   // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	r.failOnPublishError(err, "Failed to declare an exchange")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticker := time.NewTicker(time.Second / time.Duration(messageRate))
	defer ticker.Stop()

	for processedMsgs < timedMessages {
		for i := 0; i < timedMessages-processedMsgs; i++ {
			<-ticker.C
			body := "escape!"
			log.Printf("processing Message.....")
			err := ch.PublishWithContext(ctx,
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			r.failOnPublishError(err, "Failed to publish a message")

			log.Printf(" [x] Sent message: %s", body)
			log.Printf(" processed messages: %d", processedMsgs)
			processedMsgs++
		}
	}
}
