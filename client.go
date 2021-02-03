package client

import (
	"fmt"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

type Client struct {
	decisionClient DecisionClient
	userDbClient   UserDbClient
	pixelClient    PixelClient
}

type DecisionClient struct {
	path      string
	apiClient ApiClient
	networkId int
	siteId    int
}

type UserDbClient struct {
	networkId int
}

type PixelClient struct {
	apiClient ApiClient
}

func (c *DecisionClient) Get(req model.DecisionRequest, opts ...AdditionalOptions) model.DecisionResponse {
	placements := req.Placements

	for i, p := range placements {
		if p.DivName == "" {
			p.DivName = fmt.Sprintf("div%d", i)
		}

		if p.NetworkId == nil {
			p.NetworkId = &c.networkId
		}

		if p.SiteId == nil {
			p.SiteId = &c.siteId
		}

		placements[i] = p
	}

	if len(opts) > 0 {
		opt := opts[0]

		if len(opt.UserAgent) > 0 {
			c.apiClient.requestHeaders["User-Agent"] = opt.UserAgent
		} else {
			c.apiClient.requestHeaders["User-Agent"] = "OpenAPI-Generator/1.0/go"
		}

		if opt.IncludeExplanation {
			c.apiClient.requestHeaders["X-Adzerk-Explain"] = opt.ApiKey
		}
	}

	res := c.apiClient.GetDecisions(req)

	return res
}

func (c *PixelClient) Fire(opts PixelFireOptions, additionalOpts ...AdditionalOptions) model.PixelFireResponse {
	if len(additionalOpts) > 0 {
		opt := additionalOpts[0]

		if len(opt.UserAgent) > 0 {
			c.apiClient.requestHeaders["User-Agent"] = opt.UserAgent
		} else {
			c.apiClient.requestHeaders["User-Agent"] = "OpenAPI-Generator/1.0/go"
		}
	}

	res := c.apiClient.FirePixel(opts)

	return res
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
