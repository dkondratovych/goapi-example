package responses

type ResponseError struct {
	ErrorCodeId      int    `json:"errorCodeId"`
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
}
