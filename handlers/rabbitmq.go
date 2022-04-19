package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/streadway/amqp"
)

//method for repeated code
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func StartMsgBrokerConnection() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"guild_rpc", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare q")

	err = ch.Qos(
		1,     //prefetch count
		0,     //prefetch size
		false, //global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			charId, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")
			log.Printf(" [.] CharID(%d)", charId)

			char := *sendCharacterByID(charId)
			response, err := json.Marshal(char)
			failOnError(err, "Failed to convert response to json")

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			failOnError(err, "failed to publish response")

			d.Ack(false)
		}
	}()
	fmt.Println("Successfully connected to our RabbitMQ instance \n [*] - waiting for messages")
	<-forever
}

func sendCharacterByID(id int) *data.Character {
	b, i, err := data.FindChar(id)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_ = i
	return b
}
