package model

type DecisionResponse struct {
	User      User                   `json:"user"`
	Decisions map[string]Decision    `json:"decisions"`
	Explain   map[string]interface{} `json:"explain"`
}
