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
	requestHeaders := map[string]interface{}{}

	client.decisionClient = DecisionClient{
		path,
		ApiClient{path, opts.ApiKey, requestHeaders},
		opts.NetworkId,
		opts.SiteId,
	}

	client.userDbClient = UserDbClient{
		ApiClient{path, opts.ApiKey, requestHeaders},
		opts.NetworkId,
	}

	client.pixelClient = PixelClient{
		ApiClient{path, opts.ApiKey, requestHeaders},
	}

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
