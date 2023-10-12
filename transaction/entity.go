package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID int `json:"id"`
	CampaignID int `json:"campaign_id"`
	UserID int `json:"user_id"`
	User user.User
	Amount int `json:"amount"`
	Status string `json:"status"`
	Code string `json:"code"`
	PaymentURL string `json:"payment_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Campaign campaign.Campaign
}