package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

type ApiClient struct {
	basePath       string
	apiKey         string
	requestHeaders map[string]interface{}
}

func (c *ApiClient) GetDecisions(req model.DecisionRequest) model.DecisionResponse {
	url := fmt.Sprintf("%s%s", c.basePath, "/api/v2")
	values, _ := json.Marshal(req)
	res := c.request("POST", url, &values)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))

	var d model.DecisionResponse
	json.Unmarshal(body, &d)

	return d
}

func (c *ApiClient) FirePixel(opts PixelFireOptions) model.PixelFireResponse {
	if opts.RevenueOverride != nil {
		c.requestHeaders["override"] = fmt.Sprintf("%f", *opts.RevenueOverride)
	}

	if opts.AdditionalRevenue != nil {
		c.requestHeaders["additional"] = fmt.Sprintf("%f", *opts.AdditionalRevenue)
	}

	if opts.EventMultiplier != nil {
		c.requestHeaders["eventMultiplier"] = fmt.Sprintf("%d", *opts.EventMultiplier)
	}

	res := c.request("GET", opts.Url, nil)

	var location string
	locationHeaders := res.Header["Location"]
	if locationHeaders != nil && len(locationHeaders) > 0 {
		location = locationHeaders[0]
	}

	return model.PixelFireResponse{res.StatusCode, location}
}

func (c *ApiClient) request(method, url string, body *[]byte) *http.Response {
	var reqBody io.Reader

	if body != nil {
		log.Println(string(*body))
		reqBody = bytes.NewBuffer(*body)
	}

	req, _ := http.NewRequest(
		method,
		url,
		reqBody,
	)

	if c.apiKey != "" {
		req.Header.Set("X-Adzerk-ApiKey", c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")

	for k, v := range c.requestHeaders {
		req.Header.Set(k, v.(string))
	}

	client := &http.Client{
		// If CheckRedirect returns ErrUseLastResponse,
		// then the most recent response is returned with its body unclosed,
		// along with a nil error.
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, _ := client.Do(req)

	return res
}
