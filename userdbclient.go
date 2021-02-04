package kevel

import "github.com/linksports/kevel-decision-sdk-go/model"

type UserDbClient struct {
	networkId int
	apiClient ApiClient
}

func NewUserDbClient(networkId int, path, apiKey string) UserDbClient {
	return UserDbClient{
		networkId,
		NewApiClient(path, apiKey),
	}
}

func (c *UserDbClient) SetCustomProperties(userKey string, props map[string]interface{}, networkIds ...int) {
	c.apiClient.SetCustomProperties(c.NetworkId(networkIds...), userKey, props)
}

func (c *UserDbClient) AddInterest(userKey string, interest string, networkIds ...int) {
	c.apiClient.AddInterest(c.NetworkId(networkIds...), userKey, interest)
}

func (c *UserDbClient) AddRetargetingSegment(userKey string, advertiserId, retargetingSegmentId int, networkIds ...int) {
	c.apiClient.AddRetargetingSegment(c.NetworkId(networkIds...), userKey, advertiserId, retargetingSegmentId)
}

func (c *UserDbClient) Forget(userKey string, networkIds ...int) {
	c.apiClient.Forget(c.NetworkId(networkIds...), userKey)
}

func (c *UserDbClient) GdprConsent(consentRequest model.ConsentRequest, networkIds ...int) {
	c.apiClient.GdprConsent(c.NetworkId(networkIds...), consentRequest)
}

func (c *UserDbClient) IpOverride(userKey, ip string, networkIds ...int) {
	c.apiClient.IpOverride(c.NetworkId(networkIds...), userKey, ip)
}

func (c *UserDbClient) MatchUser(userKey string, partnerId, userId int, networkIds ...int) {
	c.apiClient.MatchUser(c.NetworkId(networkIds...), userKey, partnerId, userId)
}

func (c *UserDbClient) OptOut(userKey string, networkIds ...int) {
	c.apiClient.OptOut(c.NetworkId(networkIds...), userKey)
}

func (c *UserDbClient) Read(userKey string, networkIds ...int) model.UserRecord {
	c.apiClient.requestHeaders["Content-Type"] = "application/json"
	return c.apiClient.ReadUser(c.NetworkId(networkIds...), userKey)
}

func (c *UserDbClient) NetworkId(networkIds ...int) int {
	if len(networkIds) > 0 {
		return networkIds[0]
	}

	return c.networkId
}
