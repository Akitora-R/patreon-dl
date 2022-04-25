package service

import (
	"log"
	"patreon-dl/internal/app/patreon-dl/util"
	"testing"
)

func TestGetCampaign(t *testing.T) {
	util.ProxyUrl = "http://localhost:1082"
	campaign, err := GetCampaign()
	if err != nil {
		panic(err)
	}
	for _, c := range campaign {
		log.Println(c)
	}
}
