package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/doo-dev/core-mvp/internal/adapter/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Hello from core ^_^ ")

	cfg := rabbitmq.Config{
		Username: "test",
		Password: "secret",
		Host:     "localhost:5672",
		VHost:    "events",
	}

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	mqClient, err := rabbitmq.NewRabbitMq(cfg)
	if err != nil {
		panic(err)
	}

	if err := mqClient.CreateQueue("go_sdk", true, false); err != nil {
		panic(err)
	}

	if err := mqClient.CreateExchange("source_output", "topic", true, false); err != nil {
		panic(err)
	}

	if err := mqClient.CreateBinding("go_sdk", "source.output.*", "source_output"); err != nil {
		panic(err)
	}

	msgBus, err := mqClient.Consume("go_sdk", "result", false)
	if err != nil {
		fmt.Println("consumer message: ", err)
	}

	if err := mqClient.Send(context.Background(), "go_sdk", "source.*", amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Transient,
		Body:         []byte("A test message"),
	}); err != nil {
		panic(err)
	}

	go func() {
		for msg := range msgBus {
			fmt.Println("new message: ", msg)
		}
	}()

	<-done
}
