package model

type PricingData struct {
	Price      *float64 `json:"price"`
	ClearPrice *float64 `json:"clearPrice"`
	Revenue    *float64 `json:"revenue"`
	RateType   *int     `json:"rateType"`
	ECPM       *float64 `json:"eCPM"`
}
