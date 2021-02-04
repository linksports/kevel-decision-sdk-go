package model

type UserRecord struct {
	Key                 string                 `json:"key"`
	Interests           []string               `json:"interests"`
	BlockedItems        map[string][]int       `json:"blockedItems"`
	PartnerUserIds      map[string][]string    `json:"partnerUserIds"`
	RetargetingSegments map[string][]int       `json:"retargetingSegments"`
	Custom              map[string]interface{} `json:"custom"`
	Consent             map[string]bool        `json:"consent"`
	OptOut              bool                   `json:"optOut"`
	Ip                  string                 `json:"ip"`
}
