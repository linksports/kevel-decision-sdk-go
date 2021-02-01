package model

type DecisionRequest struct {
	Placements           []Placement            `json:"placements"`
	User                 User                   `json:"user"`
	Keywords             []string               `json:"keywords"`
	Url                  string                 `json:"url"`
	Referrer             string                 `json:"referrer"`
	Ip                   string                 `json:"ip"`
	BlockedCreatives     []int                  `json:"blockedCreatives"`
	IsMobile             bool                   `json:"isMobile"`
	IncludePricingData   bool                   `json:"includePricingData"`
	Notrack              bool                   `json:"notrack"`
	EnableBotFiltering   bool                   `json:"enableBotFiltering"`
	EnableUserDBIP       bool                   `json:"enableUserDBIP"`
	Consent              map[string]interface{} `json:"consent"`
	DeviceID             string                 `json:"deviceID"`
	Parallel             bool                   `json:"parallel"`
	IntendedLatitude     string                 `json:"intendedLatitude"`
	IntendedLongitude    string                 `json:"intendedLongitude"`
	IncludeMatchedPoints bool                   `json:"includeMatchedPoints"`
}

func NewDecisionRequest() DecisionRequest {
	return DecisionRequest{}
}
