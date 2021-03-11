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

func (c *ApiClient) GetDecisions(req model.DecisionRequest) (model.DecisionResponse, error) {
	body, _ := json.Marshal(req)
	c.requestHeaders["Content-Type"] = "application/json"
	res, err := c.request("POST", c.basePath, &body)

	if err != nil {
		return model.DecisionResponse{}, err
	}

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return model.DecisionResponse{}, err
	}

	var response model.DecisionResponse
	err = json.Unmarshal(body, &response)

	return response, err
}

func (c *ApiClient) FirePixel(opts PixelFireOptions) (model.PixelFireResponse, error) {
	if opts.RevenueOverride != nil {
		c.requestHeaders["override"] = fmt.Sprintf("%f", *opts.RevenueOverride)
	}

	if opts.AdditionalRevenue != nil {
		c.requestHeaders["additional"] = fmt.Sprintf("%f", *opts.AdditionalRevenue)
	}

	if opts.EventMultiplier != nil {
		c.requestHeaders["eventMultiplier"] = fmt.Sprintf("%d", *opts.EventMultiplier)
	}

	res, err := c.request("GET", opts.Url, nil)

	if err != nil {
		return model.PixelFireResponse{}, err
	}

	var location string
	locationHeaders := res.Header["Location"]
	if locationHeaders != nil && len(locationHeaders) > 0 {
		location = locationHeaders[0]
	}

	return model.PixelFireResponse{res.StatusCode, location}, nil
}

func (c *ApiClient) SetCustomProperties(networkId int, userKey string, props map[string]interface{}) error {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d/custom?%s", c.basePath, networkId, values.Encode())
	body, _ := json.Marshal(props)
	c.requestHeaders["Content-Type"] = "application/json"
	_, err := c.request("POST", urlStr, &body)
	return err
}

func (c *ApiClient) AddInterest(networkId int, userKey string, interest string) error {
	values := url.Values{
		"userKey":  {userKey},
		"interest": {interest},
	}
	urlStr := fmt.Sprintf("%s/%d/interest/i.gif?%s", c.basePath, networkId, values.Encode())
	_, err := c.request("GET", urlStr, nil)
	return err
}

func (c *ApiClient) AddRetargetingSegment(networkId int, userKey string, advertiserId, retargetingSegmentId int) error {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf(
		"%s/%d/rt/%d/%d/i.gif?%s",
		c.basePath, networkId, advertiserId, retargetingSegmentId, values.Encode())
	_, err := c.request("GET", urlStr, nil)
	return err
}

func (c *ApiClient) Forget(networkId int, userKey string) error {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d?%s", c.basePath, networkId, values.Encode())
	_, err := c.request("DELETE", urlStr, nil)
	return err
}

func (c *ApiClient) GdprConsent(networkId int, consentRequest model.ConsentRequest) error {
	urlStr := fmt.Sprintf("%s/%d/consent", c.basePath, networkId)
	body, _ := json.Marshal(consentRequest)
	c.requestHeaders["Content-Type"] = "application/json"
	_, err := c.request("POST", urlStr, &body)
	return err
}

func (c *ApiClient) IpOverride(networkId int, userKey, ip string) error {
	values := url.Values{
		"userKey": {userKey},
		"ip":      {ip},
	}
	urlStr := fmt.Sprintf("%s/%d/ip/i.gif?%s", c.basePath, networkId, values.Encode())
	_, err := c.request("GET", urlStr, nil)
	return err
}

func (c *ApiClient) MatchUser(networkId int, userKey string, partnerId, userId int) error {
	values := url.Values{
		"userKey":   {userKey},
		"partnerId": {strconv.Itoa(partnerId)},
		"userId":    {strconv.Itoa(userId)},
	}
	urlStr := fmt.Sprintf("%s/%d/sync/i.gif?%s", c.basePath, networkId, values.Encode())
	_, err := c.request("GET", urlStr, nil)
	return err
}

func (c *ApiClient) OptOut(networkId int, userKey string) error {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d/optout/i.gif?%s", c.basePath, networkId, values.Encode())
	_, err := c.request("GET", urlStr, nil)
	return err
}

func (c *ApiClient) ReadUser(networkId int, userKey string) (model.UserRecord, error) {
	values := url.Values{
		"userKey": {userKey},
	}
	urlStr := fmt.Sprintf("%s/%d/read?%s", c.basePath, networkId, values.Encode())

	res, err := c.request("GET", urlStr, nil)

	if err != nil {
		return model.UserRecord{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return model.UserRecord{}, err
	}

	var record model.UserRecord
	err = json.Unmarshal(body, &record)

	return record, err
}

func (c *ApiClient) request(method, urlStr string, body *[]byte) (*http.Response, error) {
	var reqBody io.Reader

	if body != nil {
		reqBody = bytes.NewBuffer(*body)
	}

	req, err := http.NewRequest(
		method,
		urlStr,
		reqBody,
	)

	if err != nil {
		return nil, err
	}

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
	res, err := client.Do(req)

	return res, err
}
