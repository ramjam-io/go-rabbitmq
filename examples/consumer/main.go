package main

import (
	"log"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

func main() {
	consumer, err := rabbitmq.NewConsumer(
		"amqp://user:pass@localhost",
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	err = consumer.StartConsuming(
		func(d rabbitmq.Delivery) bool {
			log.Printf("consumed: %v", string(d.Body))
			// true to ACK, false to NACK
			return true
		},
		"my_queue",
		[]string{"routing_key1", "routing_key2"},
		rabbitmq.WithConsumeOptionsConcurrency(10),
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsQuorum,
	)
	if err != nil {
		log.Fatal(err)
	}

	// block main thread so consumers run forever
	forever := make(chan struct{})
	<-forever
}
