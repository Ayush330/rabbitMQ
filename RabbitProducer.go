package main

import (
	_ "fmt"
	"context"
	"time"
	"log"
	"math/rand"
	amqp "github.com/rabbitmq/amqp091-go"
)

func producer() {
	rand.Seed(time.Now().UnixNano())
	// make connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	// declare a queue
	q, err := ch.QueueDeclare(
		"hello",
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancelFun := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFun()

	PayloadList := []string{
		// "Hello World!",
		// 1,
		// []int{1, 2, 3},
		// []string{"Hello", "World"},
		// true,
		"Hello World!",
		"Hello Ayush!",
		"Hello RabbitMQ!",
		"Hello Go!",
		"Hello Earth!",
		"Hello Mars!",
		"Hello Universe!",
		"Hello Galaxy!",
		"Hello Milky Way!",
		"Hello Solar System!",
		"Hello Planet!",
	}
	for {
		index := rand.Intn(len(PayloadList))
		body := PayloadList[index]
		if produce(ch, ctx, q.Name, body){
			log.Printf(" [x] Sent %s\n", body)
		}else{
			log.Printf(" [x] Failed to send %s\n", body)
		}
		time.Sleep(2 * time.Second)
	}

}


func produce(Channel *amqp.Channel, Context context.Context, QueueName string, payload string)bool{
	err := Channel.PublishWithContext(
		Context,
		"", // exchange
		QueueName,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(payload),},
	)
	failOnError(err, "Failed to publish a message")
	return true
}
