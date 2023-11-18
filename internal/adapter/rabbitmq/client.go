package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Host     string `koanf:"host"`
	VHost    string `koanf:"vhost"`
}

type Client struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMq(cfg Config) (Client, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.VHost))
	if err != nil {
		return Client{}, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return Client{}, err
	}

	return Client{conn: conn, ch: ch}, nil
}

func (c Client) Close() error {
	return c.Close()
}

func (c Client) CreateQueue(name string, durable, autoDelete bool) error {
	_, err := c.ch.QueueDeclare(name, durable, autoDelete, false, false, nil)

	return err
}

func (c Client) CreateExchange(name, kind string, durable, autoDelete bool) error {
	return c.ch.ExchangeDeclare(name, kind, durable, autoDelete, false, false, nil)
}

func (c Client) CreateBinding(name, binding, exchange string) error {
	return c.ch.QueueBind(name, binding, exchange, false, nil)
}

// send used to push payload to exchange with a route key
func (c Client) Send(ctx context.Context, exchange, routeKey string, options amqp.Publishing) error {
	return c.ch.PublishWithContext(ctx,
		exchange,
		routeKey,
		true,
		false,
		options,
	)
}

func (c Client) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return c.ch.Consume(queue, consumer, autoAck, false, false, false, nil)
}
