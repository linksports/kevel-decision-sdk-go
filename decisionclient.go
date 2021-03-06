package kevel

import (
	"fmt"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

type DecisionClient struct {
	networkId int
	siteId    int
	apiClient ApiClient
}

func NewDecisionClient(networkId, siteId int, path string) DecisionClient {
	return DecisionClient{
		networkId,
		siteId,
		NewApiClient(path),
	}
}

func (c *DecisionClient) Get(req model.DecisionRequest, opts ...AdditionalOptions) (model.DecisionResponse, error) {
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
		}

		if opt.IncludeExplanation {
			c.apiClient.requestHeaders["X-Adzerk-Explain"] = opt.ApiKey
		}
	}

	return c.apiClient.GetDecisions(req)
}
