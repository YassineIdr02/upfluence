package types

type Campaign struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Influencer string `json:"influencer"`
	Budget     int    `json:"budget"`
	Status     string `json:"status"`
}

const RabbitMQURL = "amqp://guest:guest@localhost:5672/"

const (
	CampaignStatusActive   = "active"
	CampaignStatusInactive = "inactive"
)
