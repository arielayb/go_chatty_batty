package app

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockAmqDialer struct {
	RmqConn     *amqp.Connection
	ReturnError error
}

func (m *mockAmqDialer) Dial(url string) (*amqp.Connection, error) {
	return m.RmqConn, m.ReturnError
}

func TestRmqConnect(t *testing.T) {
	mockAmqDialer := &mockAmqDialer{
		RmqConn:     &amqp.Connection{},
		ReturnError: nil,
	}
	rmqClient := NewRmqClient(mockAmqDialer)
	if err := rmqClient.RmqConnect("testme"); err != nil {
		t.Errorf("Error: %v", err.Error())
	}

	assert.NotNil(t, rmqClient, "the connection shouldn't be null")
}
