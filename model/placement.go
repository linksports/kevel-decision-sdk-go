package model

type Placement struct {
	DivName         string                 `json:"divName"`
	NetworkId       *int                   `json:"networkId"`
	SiteId          *int                   `json:"siteId"`
	AdTypes         []int                  `json:"adTypes"`
	ZoneIds         []int                  `json:"zoneIds"`
	CampaignId      *int                   `json:"campaignId"`
	FlightId        *int                   `json:"flightId"`
	AdId            *int                   `json:"adId"`
	ClickUrl        string                 `json:"clickUrl"`
	Properties      map[string]interface{} `json:"properties"`
	EventIds        []int                  `json:"eventIds"`
	Overrides       map[string]interface{} `json:"overrides"`
	ContentKeys     map[string]int         `json:"contentKeys"`
	Count           *int                   `json:"count"`
	Proportionality bool                   `json:"proportionality"`
	EcpmPartition   string                 `json:"expmPartition"`
	EventMultiplier *int                   `json:"eventMultiplier"`
	SkipSelection   bool                   `json:"skipSelection"`
}

func NewPlacement() Placement {
	return Placement{}
}
