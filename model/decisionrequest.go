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
	IntendedLatitude     *float64               `json:"intendedLatitude"`
	IntendedLongitude    *float64               `json:"intendedLongitude"`
	Radius               *float64               `json:"radius"`
	IncludeMatchedPoints bool                   `json:"includeMatchedPoints"`
}

func NewDecisionRequest() DecisionRequest {
	return DecisionRequest{}
}
