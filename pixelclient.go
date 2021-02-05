package kevel

import "github.com/linksports/kevel-decision-sdk-go/model"

type PixelClient struct {
	apiClient ApiClient
}

func NewPixelClient() PixelClient {
	return PixelClient{ApiClient{}}
}

func (c *PixelClient) Fire(opts PixelFireOptions, additionalOpts ...AdditionalOptions) (model.PixelFireResponse, error) {
	if len(additionalOpts) > 0 {
		opt := additionalOpts[0]

		if len(opt.UserAgent) > 0 {
			c.apiClient.requestHeaders["User-Agent"] = opt.UserAgent
		}
	}

	return c.apiClient.FirePixel(opts)
}
