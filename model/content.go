package model

type Content struct {
	Type           string                 `json:"type"`
	Template       string                 `json:"template"`
	CustomTemplate string                 `json:"customTemplate"`
	Data           map[string]interface{} `json:"data"`
	Body           string                 `json:"body"`
}
