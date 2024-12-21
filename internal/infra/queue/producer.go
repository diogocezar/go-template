package queue

import (
	"fmt"
	"go-template/internal/config"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Queue   amqp091.Queue
	Channel *amqp091.Channel
}

func MakeProducer(envs *config.Envs, queueName string, body string) *Producer {

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
		log.Fatalf("Error trying to declarate Queue: %v", err)
	}

	return &Producer{
		Queue:   queue,
		Channel: ch,
	}
}

func (p *Producer) Publish(body string, exchange string) {
	ch := p.Channel

	err := ch.Publish(
		exchange,     // exchange
		p.Queue.Name, // routing key
		false,        // obrigatório
		false,        // imediato
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Fatalf("Error trying to publish message: %v", err)
	}

	log.Printf("Message sent: %s", body)
}
