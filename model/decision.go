package model

type Decision struct {
	AdId          *int           `json:"adId"`
	CreativeId    *int           `json:"creativeId"`
	FlightId      *int           `json:"flightId"`
	CampaignId    *int           `json:"campaignId"`
	PriorityId    *int           `json:"priorityId"`
	ClickUrl      string         `json:"clickUrl"`
	Contents      []Content      `json:"contents"`
	ImpressionUrl string         `json:"impressionUrl"`
	Events        []Event        `json:"events"`
	MatchedPoints []MatchedPoint `json:"matchedPoints"`
	Pricing       PricingData    `json:"pricing"`
}
