package entity

import (
	"time"
)

type PledgeResp struct {
	Data []struct {
		Attributes struct {
			AmountCents        int       `json:"amount_cents"`
			Cadence            int       `json:"cadence"`
			CreatedAt          time.Time `json:"created_at"`
			Currency           string    `json:"currency"`
			HasShippingAddress bool      `json:"has_shipping_address"`
			IsPaused           bool      `json:"is_paused"`
			PledgeCapCents     int       `json:"pledge_cap_cents"`
			Status             string    `json:"status"`
		} `json:"attributes"`
		Id            string `json:"id"`
		Relationships struct {
			Address struct {
				Data interface{} `json:"data"`
			} `json:"address"`
			Campaign struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"campaign"`
			MostRecentPledgeChargeTxn struct {
				Data *struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links,omitempty"`
			} `json:"most_recent_pledge_charge_txn"`
			Reward struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reward"`
		} `json:"relationships"`
		Type string `json:"type"`
	} `json:"data"`
	Included []struct {
		Attributes struct {
			Description      *string     `json:"description,omitempty"`
			RequiresShipping bool        `json:"requires_shipping,omitempty"`
			UnpublishedAt    interface{} `json:"unpublished_at"`
			AvatarPhotoUrl   string      `json:"avatar_photo_url,omitempty"`
			CoverPhotoUrl    string      `json:"cover_photo_url,omitempty"`
			IsMonthly        bool        `json:"is_monthly,omitempty"`
			IsNonProfit      bool        `json:"is_non_profit,omitempty"`
			Name             string      `json:"name,omitempty"`
			PayPerName       string      `json:"pay_per_name,omitempty"`
			PledgeUrl        string      `json:"pledge_url,omitempty"`
			PublishedAt      time.Time   `json:"published_at,omitempty"`
			Url              string      `json:"url,omitempty"`
			FailedAt         interface{} `json:"failed_at"`
			SucceededAt      time.Time   `json:"succeeded_at,omitempty"`
			EndedAt          interface{} `json:"ended_at"`
			IsEnded          bool        `json:"is_ended,omitempty"`
			IsPublished      bool        `json:"is_published,omitempty"`
			ItemType         string      `json:"item_type,omitempty"`
			Title            string      `json:"title,omitempty"`
		} `json:"attributes"`
		Id            string `json:"id"`
		Relationships struct {
			Items struct {
				Data []struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"items,omitempty"`
			MerchCustomVariants struct {
				Data []interface{} `json:"data"`
			} `json:"merch_custom_variants,omitempty"`
			RewardItemConfiguration struct {
				Data interface{} `json:"data"`
			} `json:"reward_item_configuration,omitempty"`
		} `json:"relationships,omitempty"`
		Type string `json:"type"`
	} `json:"included"`
}

type Campaign struct {
	Id   string
	Name string
	Url  string
}
