package kevel

import "github.com/linksports/kevel-decision-sdk-go/model"

type UserDbClient struct {
	apiClient ApiClient
	networkId int
}

func (c *UserDbClient) Read(userKey string, networkIds ...int) model.UserRecord {
	networkId := c.networkId

	if len(networkIds) > 0 {
		networkId = networkIds[0]
	}

	record := c.apiClient.ReadUserDb(userKey, networkId)

	return record
}
