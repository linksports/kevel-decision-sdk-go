package model

type ConsentRequest struct {
	UserKey string                 `json:"userKey"`
	Consent map[string]interface{} `json:"consent"`
}

func NewConsentRequest(userKey string, consent map[string]interface{}) ConsentRequest {
	return ConsentRequest{userKey, consent}
}
