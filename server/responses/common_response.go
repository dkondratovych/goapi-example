package responses

type CommonResponse struct {
	Data     interface{}            `json:"data"`
	Metadata map[string]interface{} `json:"metadata"`
}
