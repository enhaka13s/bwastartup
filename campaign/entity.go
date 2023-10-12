package campaign

import (
	"bwastartup/user"
	"time"
)

type Campaign struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
	ShortDescription string `json:"short_description"`
	Description string `json:"description"`
	Perks string `json:"perks"`
	BackerCount int `json:"backer_count"`
	GoalAmount int `json:"goal_amount"`
	CurrentAmount int `json:"current_amount"`
	Slug string `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CampaignImages []CampaignImage `json:"campaign_images"`
	User user.User
}

type CampaignImage struct {
	ID int
	CampaignID int
	FileName string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time	
}