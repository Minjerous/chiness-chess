package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Claim struct {
	User string
	Host string
	Port string
	PW   string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func SendMessage(claim *Claim, Uid int64) {

	conn, err := amqp.Dial("amqp://" + claim.User + ":" + claim.PW + "@" + claim.Host + ":" + claim.Port + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"room_math", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := fmt.Sprintln(Uid, "加入对局")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}
