package handlers

import (
	"encoding/json"
	"log"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/streadway/amqp"
)

//method for repeated code
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendCharsByGID(gId string) {
	//connect to  rabbit mq server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//declare channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declare queue
	q, err := ch.QueueDeclare(
		"guild_rpc", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	//get character by guildID
	char, err := data.GetCharactersByGuild(string(gId))
	if char == nil {
		log.Print("no characters with that guild id found", err)

	}

	response, err := json.Marshal(char)
	failOnError(err, "Failed to convert response to json")

	corrId := "getall"
	//publish message with characters by guildID
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			Body:          response,
		})
	failOnError(err, "failed to publish response")
}
