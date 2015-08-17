package config

type Server struct {
	Port string `json:"port"`
	HttpLogFilePath string `json:"http_log_file"`
	JwtSecret string `json:"jwt_secret"`
}