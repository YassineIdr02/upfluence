package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

	campaigns := []types.Campaign{
		{ID: 1, Name: "Back to School", Influencer: "John Doe", Budget: 1000, Status: types.CampaignStatusActive},
		{ID: 2, Name: "Holiday Sale", Influencer: "Jane Smith", Budget: 2000, Status: types.CampaignStatusActive},
		{ID: 3, Name: "New Product Launch", Influencer: "Alice Johnson", Budget: 3000, Status: types.CampaignStatusInactive},
	}

	for _, campaign := range campaigns {
		body, err := json.Marshal(campaign)
		failOnError(err, "Failed to marshal campaign")

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		failOnError(err, "Failed to publish message")

		fmt.Printf("Sent campaign notification: %s\n", campaign.Name)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All campaign messages sent!")
}
