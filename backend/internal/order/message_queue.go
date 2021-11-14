package order

import "github.com/streadway/amqp"

type messageQueue struct {
	conn *amqp.Connection
}

func NewOrderMessageQueue(conn *amqp.Connection) OrderMessageQueue {
	return &messageQueue{conn: conn}
}
