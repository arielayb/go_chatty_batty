package app

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRmqpClient struct {
}

func (m mockRmqpClient) Dial(amqpDial AmqpDial, url string) (*amqp.Connection, error) {

	conn, err := AmqpDial.Dial(amqpDial, url)
	if err != nil {
		fmt.Errorf("cannot connect!")
	}
	return conn, nil
}

func TestingRMQConn(t *testing.T) {
	mockRmqpClient := mockRmqpClient{}
	conn, err := AmqpDial.Dial(mockRmqpClient, "amqp://testing.com")
	assert.NotNil(t, conn, "the connection shouldn't be null")
}
