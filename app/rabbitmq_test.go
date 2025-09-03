package app

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRmqClient struct {
}

func (m mockRmqClient) Dial(amqpDial AmqpDial, url string) (*amqp.Connection, error) {
	return &amqp.Connection{}, nil
}

func TestingRMQConn(t *testing.T) {
	mockRmqClient := mockRmqClient{}
	conn, err := rmqConnect("ampq://testme")
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	assert.NotNil(t, conn, "the connection shouldn't be null")
}
