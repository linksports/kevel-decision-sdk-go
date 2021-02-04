package kevel

import "github.com/linksports/kevel-decision-sdk-go/model"

type PixelClient struct {
	apiClient ApiClient
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
