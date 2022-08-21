package rabbitmq

import (
	"context"
	"fmt"

	"github.com/briand787b/ncx/src/svc/pkg/cserr"
	"github.com/briand787b/ncx/src/svc/pkg/cslog"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

// NewConnection connects to RabbitMQ and initializes configuration in the broker.
// Please defer the closing of the returned connection
func NewConnection(ctx context.Context, l cslog.Logger, host string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(host)
	if err != nil {
		return nil, cserr.NewErrInternalWithMsg(err, "could not dial host to create connection")
	}

	if err := createQueueBoundDirectExchange(conn, TicketDiscoveredExchange, TicketDiscoveredQueue); err != nil {
		return nil, errors.Wrapf(err, "could not create queue-bound direct exchange %s", TicketDiscoveredExchange)
	}

	if err := createQueueBoundDirectExchange(conn, LocationCreatedExchange, LocationCreatedQueue); err != nil {
		return nil, errors.Wrapf(err, "could not create queue-bound direct exchange %s", LocationCreatedExchange)
	}

	if err := createQueueBoundDirectExchange(conn, LocationDiscoveredExchange, LocationDiscoveredQueue); err != nil {
		return nil, errors.Wrapf(err, "could not create queue-bound direct exchange %s", LocationDiscoveredExchange)
	}

	go func() {
		for err := range conn.NotifyClose(make(chan *amqp.Error)) {
			l.Error(ctx, "error on the amqp connection", "error", err.Error())
		}
	}()

	return conn, nil
}

func createQueueBoundDirectExchange(cn *amqp.Connection, exchange string, queue string) error {
	ch, err := cn.Channel()
	if err != nil {
		return cserr.NewErrInternalWithMsg(err, "could not create channel")
	}

	defer ch.Close()

	if err := ch.ExchangeDeclare(
		exchange,            // name
		amqp.ExchangeDirect, // kind
		false,               // durable
		false,               // autoDelete
		false,               // internal
		false,               // noWait
		nil,                 // args
	); err != nil {
		return cserr.NewErrInternalWithMsg(err, fmt.Sprintf("could not declare exchange '%s'", exchange))
	}

	if _, err = ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	); err != nil {
		return cserr.NewErrInternalWithMsg(err, fmt.Sprintf("could not declare queue '%s'", queue))
	}

	if err := ch.QueueBind(
		queue,    // queueName
		"",       // routingKey
		exchange, // exchangeName
		false,    // noWait
		nil,      // args
	); err != nil {
		return cserr.NewErrInternalWithMsg(err, fmt.Sprintf("could not bind queue '%s' to exchange '%s'", queue, exchange))
	}

	return nil
}
