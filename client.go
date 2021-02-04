package kevel

import "fmt"

type Client struct {
	decisionClient DecisionClient
	userDbClient   UserDbClient
	pixelClient    PixelClient
}

func NewClient(opts ClientOptions) Client {
	protocol := opts.Protocol
	if len(protocol) == 0 {
		protocol = "https"
	}

	host := opts.Host
	if len(host) == 0 {
		host = fmt.Sprintf("e-%d.adzerk.net", opts.NetworkId)
	}

	path := fmt.Sprintf("%s://%s", protocol, host)

	client := Client{}

	client.decisionClient = NewDecisionClient(
		opts.NetworkId,
		opts.SiteId,
		fmt.Sprintf("%s/api/v2", path),
	)

	client.userDbClient = NewUserDbClient(
		opts.NetworkId,
		fmt.Sprintf("%s/udb", path),
		opts.ApiKey,
	)

	client.pixelClient = NewPixelClient()

	return client
}

func (c *Client) Decisions() DecisionClient {
	return c.decisionClient
}

func (c *Client) UserDb() UserDbClient {
	return c.userDbClient
}

func (c *Client) Pixels() PixelClient {
	return c.pixelClient
}
