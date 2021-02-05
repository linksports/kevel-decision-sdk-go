package kevel

import (
	"strings"
	"testing"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

const (
	NetworkId  = 23
	SiteId     = 667480
	UserKey    = "abc"
	AdId       = 2104402
	CreativeId = 1773302
	FlightId   = 2583477
	CampaignId = 502103
	PriorityId = 99645
)

func TestDecisionApi(t *testing.T) {
	opts := NewClientOptions(NetworkId)
	opts.SiteId = SiteId
	client := NewClient(opts)

	placement := model.NewPlacement()
	placement.AdTypes = []int{5}

	request := model.NewDecisionRequest()
	request.Placements = []model.Placement{placement}
	request.Keywords = []string{"keyword1", "keyword2"}
	request.User = model.NewUser(UserKey)

	decisions := client.Decisions()
	response, err := decisions.Get(request)

	if err != nil {
		t.Errorf("Error returned when fetching an ad decision: %s", err.Error())
	}

	if response.User.Key != UserKey {
		t.Errorf("Invalid userKey, got: %s, want: %s", response.User.Key, UserKey)
	}

	decision := response.Decisions["div0"]

	if *decision.AdId != AdId {
		t.Errorf("Invalid adId, got: %d, want: %d", *decision.AdId, AdId)
	}

	if *decision.CreativeId != CreativeId {
		t.Errorf("Invalid creativeId, got: %d, want: %d", *decision.CreativeId, CreativeId)
	}

	if *decision.FlightId != FlightId {
		t.Errorf("Invalid flightId, got: %d, want: %d", *decision.FlightId, FlightId)
	}

	if *decision.CampaignId != CampaignId {
		t.Errorf("Invalid campaignId, got: %d, want: %d", *decision.CampaignId, CampaignId)
	}

	if *decision.PriorityId != PriorityId {
		t.Errorf("Invalid priorityId, got: %d, want: %d", *decision.PriorityId, PriorityId)
	}

	if !strings.HasPrefix(decision.ClickUrl, "https://e-23.adzerk.net/r") {
		t.Errorf("Invalid clickUrl, does not start with: %s", "https://e-23.adzerk.net/r")
	}

	if !strings.HasPrefix(decision.ImpressionUrl, "https://e-23.adzerk.net/i.gif") {
		t.Errorf("Invalid impressionUrl, does not start with: %s", "https://e-23.adzerk.net/i.gif")
	}

	content := decision.Contents[0]

	if content.Type != "html" {
		t.Errorf("Invalid content type, got: %s, want: %s", content.Type, "html")
	}

	if content.Template != "image" {
		t.Errorf("Invalid content template, got: %s, want: %s", content.Template, "image")
	}

	height := int(content.Data["height"].(float64))

	if height != 250 {
		t.Errorf("Invalid height, got: %d, want: %d", height, 250)
	}

	width := int(content.Data["width"].(float64))

	if width != 300 {
		t.Errorf("Invalid width, got: %d, want: %d", width, 300)
	}

	if !strings.HasPrefix(content.Data["imageUrl"].(string), "https://s.zkcdn.net/") {
		t.Errorf("Invalid imageUrl, does not start with: %s", "https://s.zkcdn.net/")
	}

	if content.Data["title"] != "" {
		t.Errorf("Invalid title, got: %s, want: %s", content.Data["title"], "")
	}

	if content.Data["fileName"] != "fdb7324f69c6420db2947dba83e15868.png" {
		t.Errorf("Invalid fileName, got: %s, want: %s", content.Data["fileName"], "fdb7324f69c6420db2947dba83e15868.png")
	}

	pixels := client.Pixels()
	impResponse, err := pixels.Fire(NewPixelFireOptions(decision.ImpressionUrl))

	if err != nil {
		t.Errorf("Error returned when recording an impression: %s", err.Error())
	}

	if impResponse.StatusCode != 200 {
		t.Errorf("Invalid statusCode, got: %d, want: %d", impResponse.StatusCode, 200)
	}

	clickResponse, err := pixels.Fire(NewPixelFireOptions(decision.ClickUrl))

	if err != nil {
		t.Errorf("Error returned when recording a click: %s", err.Error())
	}

	if clickResponse.StatusCode != 302 {
		t.Errorf("Invalid statusCode, got: %d, want: %d", clickResponse.StatusCode, 302)
	}

	if clickResponse.Location != "https://kevel.co" {
		t.Errorf("Invalid location, got: %s, want: %s", clickResponse.Location, "https://kevel.co")

	}
}

func TestUserDb(t *testing.T) {
	opts := NewClientOptions(NetworkId)
	client := NewClient(opts)

	userDb := client.UserDb()
	record, err := userDb.Read(UserKey)

	if err != nil {
		t.Errorf("Error returned when reading user record: %s", err.Error())
	}

	if record.Key != UserKey {
		t.Errorf("Invalid userKey, got: %s, want: %s", record.Key, UserKey)
	}

	props := map[string]interface{}{
		"favoriteColor":  "blue",
		"favoriteNumber": 42,
		"favoriteFoods":  []string{"strawberries", "chocolate"},
	}

	err = userDb.SetCustomProperties(UserKey, props)

	if err != nil {
		t.Errorf("Error returned when setting custom properties: %s", err.Error())
	}
}
