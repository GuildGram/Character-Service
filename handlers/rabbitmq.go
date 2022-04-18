package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/GuildGram/Character-Service.git/data"
	"github.com/streadway/amqp"
)

func StartMsgBrokerConnection() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Success Connection RabbitmQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQ", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	var body data.Character
	body = *sendCharacterByID(1)
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = ch.Publish(
		"",      // exchange
		"TestQ", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("successfully published msg to q")
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
