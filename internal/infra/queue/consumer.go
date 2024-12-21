package queue

import (
	"fmt"
	"go-template/internal/config"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func MakeConsumer(envs *config.Envs, queueName string, handler func([]byte) error) {

	USER := envs.QUEUE_USER
	PASSWORD := envs.QUEUE_PASSWORD
	HOST := envs.QUEUE_HOST
	PORT := envs.QUEUE_PORT

	conn, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", USER, PASSWORD, HOST, PORT))
	if err != nil {
		log.Fatalf("Error trying to connect on Queue: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error trying to open channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName, // name of queue
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // extra args
	)

	if err != nil {
		log.Fatalf("Error trying to declare queue : %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // name of queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no wait
		false,      // no extra args
		nil,        // extra args
	)
	if err != nil {
		log.Fatalf("Error trying to register consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)
			handler(d.Body)
		}
	}()
	<-forever
}
