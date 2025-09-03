package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRmqClient struct {
	createRmqError error
}

func (m mockRmqClient) Dial(url string) (*amqp.Connection, error) {
	return &amqp.Connection{}, m.createRmqError
}

func TestRmqConnect(t *testing.T) {
	mockRmqClient := mockRmqClient{}
	conn, err := RmqConnect(mockRmqClient, "testme")
	if err != nil {
		t.Errorf("Error: %v", err.Error())
	}
	assert.NotNil(t, conn, "the connection shouldn't be null")
}
