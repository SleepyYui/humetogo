package hume

import (
	"github.com/sleepyyui/humetogo/api"
)

func New(apiKey string, apiSecret string) *api.HumeApi {
	return &api.HumeApi{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
		BaseUrl:   "https://api.hume.ai/v0",
		Debug:     false,
		Version:   "v1",
	}
}
