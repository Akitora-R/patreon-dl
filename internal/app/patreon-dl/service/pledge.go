package service

import (
	"encoding/json"
	"errors"
	"net/url"
	"patreon-dl/internal/app/patreon-dl/entity"
	"patreon-dl/internal/app/patreon-dl/util"
)

var (
	pledgeQuery = map[string][]string{
		"include":                       {"address,campaign,reward.items,most_recent_pledge_charge_txn,reward.items.reward_item_configuration,reward.items.merch_custom_variants,reward.items.merch_custom_variants.item,reward.items.merch_custom_variants.merch_product_variant"},
		"fields[address]":               {"id,addressee,line_1,line_2,city,state,postal_code,country,phone_number"},
		"fields[campaign]":              {"avatar_photo_url,cover_photo_url,is_monthly,is_non_profit,name,pay_per_name,pledge_url,published_at,url"},
		"fields[user]":                  {"thumb_url,url,full_name"},
		"fields[pledge]":                {"amount_cents,currency,pledge_cap_cents,cadence,created_at,has_shipping_address,is_paused,status"},
		"fields[reward]":                {"description,requires_shipping,unpublished_at"},
		"fields[reward-item]":           {"id,title,description,requires_shipping,item_type,is_published,is_ended,ended_at,reward_item_configuration"},
		"fields[merch-custom-variant]":  {"id,item_id"},
		"fields[merch-product-variant]": {"id,color,size_code"},
		"fields[txn]":                   {"succeeded_at,failed_at"},
		"json-api-use-default-includes": {"false"},
		"json-api-version":              {"1.0"},
	}
)

func getPledge() (entity.PledgeResp, error) {
	var re entity.PledgeResp
	params := url.Values{}
	for k, v := range pledgeQuery {
		params[k] = v
	}
	s := "https://www.patreon.com/api/pledges?" + params.Encode()
	b, r, err := util.NewGetRequest(s).SetHeader(header).SetOptionalProxy().Exec()
	if err != nil {
		return re, err
	}
	if r.StatusCode != 200 {
		return re, errors.New(r.Status)
	}
	if err = json.Unmarshal(b, &re); err != nil {
		return re, err
	}
	return re, nil
}

func GetCampaign() ([]entity.Campaign, error) {
	var l []entity.Campaign
	pledge, err := getPledge()
	if err != nil {
		return nil, err
	}
	for _, d := range pledge.Data {
		cId := d.Relationships.Campaign.Data.Id
		for _, inc := range pledge.Included {
			if inc.Id == cId {
				l = append(l, entity.Campaign{
					Id:   cId,
					Name: inc.Attributes.Name,
					Url:  inc.Attributes.Url,
				})
				break
			}
		}
	}
	return l, nil
}
