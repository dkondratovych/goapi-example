package handlers

type OptionsHandler struct{}

type Method struct {
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
	Example     interface{} `json:"example"`
}

type Parameters []Parameter
type Parameter map[string]interface{}
type ParameterOption map[string]interface{}
