package config

type Facebook struct {
	AppId       string `json:"appId"`
	Secret      string `json:"secret"`
	CallbackURL string `json:"callbackURL"`
}
