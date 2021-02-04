package kevel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

type ApiClient struct {
	basePath       string
	apiKey         string
	requestHeaders map[string]interface{}
}

func NewApiClient(path string, apiKey ...string) ApiClient {
	apiClient := ApiClient{}
	apiClient.basePath = path

	if len(apiKey) > 0 {
		apiClient.apiKey = apiKey[0]
	}

	apiClient.requestHeaders = map[string]interface{}{
		"User-Agent": "OpenAPI-Generator/1.0/go",
	}

	return apiClient
}

func (c *ApiClient) GetDecisions(req model.DecisionRequest) model.DecisionResponse {
	body, _ := json.Marshal(req)
	res := c.request("POST", c.basePath, &body)

	defer res.Body.Close()

	body, _ = ioutil.ReadAll(res.Body)

	var response model.DecisionResponse
	json.Unmarshal(body, &response)

	return response
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

func (c *ApiClient) SetCustomProperties(networkId int, userKey string, props map[string]interface{}) {
	urlStr := fmt.Sprintf("%s/%d/custom", c.basePath, networkId)
	body, _ := json.Marshal(props)
	c.request("POST", urlStr, &body)
}

func (c *ApiClient) AddInterest(networkId int, userKey string, interest string) {
	values := url.Values{
		"userKey":  {userKey},
		"interest": {interest},
	}
	urlStr := fmt.Sprintf("%s/%d/interest/i.gif?%s", c.basePath, networkId, values.Encode())
	c.request("GET", urlStr, nil)
}

func (c *ApiClient) AddRetargetingSegment(networkId int, userKey string, advertiserId, retargetingSegmentId int) {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf(
		"%s/%d/rt/%d/%d/i.gif?%s",
		c.basePath, networkId, advertiserId, retargetingSegmentId, values.Encode())
	c.request("GET", urlStr, nil)
}

func (c *ApiClient) Forget(networkId int, userKey string) {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d?%s", c.basePath, networkId, values.Encode())

	c.request("DELETE", urlStr, nil)
}

func (c *ApiClient) GdprConsent(networkId int, consentRequest model.ConsentRequest) {
	urlStr := fmt.Sprintf("%s/%d/consent", c.basePath, networkId)
	body, _ := json.Marshal(consentRequest)
	c.request("POST", urlStr, &body)

}

func (c *ApiClient) IpOverride(networkId int, userKey, ip string) {
	values := url.Values{
		"userKey": {userKey},
		"ip":      {ip},
	}
	urlStr := fmt.Sprintf("%s/%d/ip/i.gif?%s", c.basePath, networkId, values.Encode())
	c.request("GET", urlStr, nil)
}

func (c *ApiClient) MatchUser(networkId int, userKey string, partnerId, userId int) {
	values := url.Values{
		"userKey":   {userKey},
		"partnerId": {strconv.Itoa(partnerId)},
		"userId":    {strconv.Itoa(userId)},
	}
	urlStr := fmt.Sprintf("%s/%d/sync/i.gif?%s", c.basePath, networkId, values.Encode())
	c.request("GET", urlStr, nil)
}

func (c *ApiClient) OptOut(networkId int, userKey string) {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d/optout/i.gif?%s", c.basePath, networkId, values.Encode())
	c.request("GET", urlStr, nil)
}

func (c *ApiClient) ReadUser(networkId int, userKey string) model.UserRecord {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d/read?%s", c.basePath, networkId, values.Encode())

	res := c.request("GET", urlStr, nil)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var record model.UserRecord
	json.Unmarshal(body, &record)

	return record
}

func (c *ApiClient) request(method, urlStr string, body *[]byte) *http.Response {
	var reqBody io.Reader

	if body != nil {
		reqBody = bytes.NewBuffer(*body)
	}

	req, _ := http.NewRequest(
		method,
		urlStr,
		reqBody,
	)

	if c.apiKey != "" {
		req.Header.Set("X-Adzerk-ApiKey", c.apiKey)
	}

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
