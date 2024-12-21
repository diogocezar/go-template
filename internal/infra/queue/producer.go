package queue

import (
	"fmt"
	"go-template/internal/config"
	"log"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type Producer struct {
	Connection *amqp091.Connection
	Channel    *amqp091.Channel
	mutex      sync.Mutex
}

func MakeProducer(envs *config.Envs) (*Producer, error) {
	USER := envs.QUEUE_USER
	PASSWORD := envs.QUEUE_PASSWORD
	HOST := envs.QUEUE_HOST
	PORT := envs.QUEUE_PORT

	conn, err := amqp091.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", USER, PASSWORD, HOST, PORT))
	if err != nil {
		log.Fatalf("Error trying to connect on Queue: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Producer{
		Connection: conn,
		Channel:    ch,
	}, nil
}

func CloseProducer(p *Producer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.Channel != nil {
		p.Channel.Close()
	}
	if p.Connection != nil {
		p.Connection.Close()
	}
}

func Publish(queue string, body string, p *Producer) error {

	err := p.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		log.Fatalf("Error trying to publish message: %v", err)
	}

	log.Printf("Message published: %s", body)
	return nil
}
