package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	params, _ := json.Marshal(req)
	body := c.request("POST", "/api/v2", params)

	var res model.DecisionResponse
	json.Unmarshal(body, &res)

	return res
}

func (c *ApiClient) request(method, path string, params []byte) []byte {
	log.Println(string(params))

	req, _ := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.basePath, path),
		bytes.NewBuffer(params),
	)

	if c.apiKey == "" {
		req.Header.Set("X-Adzerk-ApiKey", c.apiKey)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	log.Println(string(body))

	return body
}
