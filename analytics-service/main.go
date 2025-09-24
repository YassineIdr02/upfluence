package main

import (
	"encoding/json"
	"fmt"
	"log"

	"upfluence-mini-project/types"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(types.RabbitMQURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"campaign_notifications",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	fmt.Println("Waiting for campaign messages...")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var campaign types.Campaign
			err := json.Unmarshal(d.Body, &campaign)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}
			fmt.Printf("Processing campaign: ID=%d, Name=%s\n", campaign.ID, campaign.Name)
			// ...
		}
	}()

	<-forever
}
