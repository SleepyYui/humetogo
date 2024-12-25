package api

type HumeApi struct {
	ApiKey    string
	ApiSecret string

	BaseUrl string
	Debug   bool
	Version string
}
