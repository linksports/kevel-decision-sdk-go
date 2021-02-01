package client

import (
	"strings"
	"testing"

	"github.com/linksports/kevel-decision-sdk-go/model"
)

func TestDecisionApi(t *testing.T) {
	opts := NewClientOptions(23) // networkId
	opts.SiteId = 667480         // siteId
	client := NewClient(opts)

	placement := model.NewPlacement()
	placement.AdTypes = []int{5}

	req := model.NewDecisionRequest()
	req.Placements = []model.Placement{placement}
	req.User = model.NewUser("abc")

	decisions := client.Decisions()
	res := decisions.Get(req)

	if res.User.Key != "abc" {
		t.Errorf("Invalid userKey, got: %s, want: %s", res.User.Key, "abc")
	}

	decision := res.Decisions["div0"]

	if *decision.AdId != 2104402 {
		t.Errorf("Invalid adId, got: %d, want: %d", *decision.AdId, 2104402)
	}

	if *decision.CreativeId != 1773302 {
		t.Errorf("Invalid creativeId, got: %d, want: %d", *decision.CreativeId, 1773302)
	}

	if *decision.FlightId != 2583477 {
		t.Errorf("Invalid flightId, got: %d, want: %d", *decision.FlightId, 2583477)
	}

	if *decision.CampaignId != 502103 {
		t.Errorf("Invalid campaignId, got: %d, want: %d", *decision.CampaignId, 502103)
	}

	if *decision.PriorityId != 99645 {
		t.Errorf("Invalid priorityId, got: %d, want: %d", *decision.PriorityId, 99645)
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
}
