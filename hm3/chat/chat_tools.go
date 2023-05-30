package chat

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(url, name, username string, msgs <-chan string) {
	conn, err := amqp.Dial(url)
	if err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return
	}
	defer ch.Close()

	for msg := range msgs {
		ctx := context.Background()
        
		err := ch.PublishWithContext(ctx, name, "", false, false, amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("[%v] %v", username, msg)),
		})
		if err != nil {
			failOnError(err, "Failed to publish message")
			return
		}
	}
}

func Consume(url, exchange_name, username string, s *StdoutWriter) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange_name, // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		failOnError(err, "Failed to declare exchange")
		return
	}
	q, err := ch.QueueDeclare(
		"", // name
		false,   // durable
		false,   // delete when unused
		true,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	log.Printf("Queue name %v", q.Name)
	if err != nil {
		failOnError(err, "Failed to declare queue")
		return
	}
	err = ch.QueueBind(q.Name, "", exchange_name, false, nil)
	if err != nil {
		failOnError(err, "Failed to bind queue")
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		username,     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		failOnError(err, "Failed to bind queue")
		return
	}

	for msg := range msgs {
		s.PrintUserMessage(string(msg.Body))
	}
	log.Println("End of publish.")
}
