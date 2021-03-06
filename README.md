# Kevel Go Decision SDK

Go Software Development Kit for Kevel Decision & UserDB APIs

## Installation

Install the go library
```
go get github.com/linksports/kevel-decision-sdk-go
```

## Examples

### API Credentials & Required IDs

- Network ID: Log into [Kevel UI](https://app.kevel.co/) & use the "circle-i" help menu in upper right corner to find Network ID. Required for all SDK operations.
- Site ID: Go to [Manage Sites page](https://app.kevel.co/#!/sites/) to find site IDs. Required when fetching an ad decision.
- Ad Type ID: Go to [Ad Sizes page](https://app.kevel.co/#!/ad-sizes/) to find Ad Type IDs. Required when fetching an ad decision.
- API Key: Go to [API Keys page](https://app.kevel.co/#!/api-keys/) find active API keys. Required when writing to UserDB.
- User Key: UserDB IDs are [specified or generated for each user](https://dev.kevel.co/reference/userdb#passing-the-userkey).

### Fetching an Ad Decision

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/linksports/kevel-decision-sdk-go"
	"github.com/linksports/kevel-decision-sdk-go/model"
)

func main() {
	// Demo network, site, and ad type IDs; find your own via the Kevel UI!
	opts := kevel.NewClientOptions(23)
	opts.SiteId = 667480
	client := kevel.NewClient(opts)

	placement := model.NewPlacement()
	placement.AdTypes = []int{5}

	request := model.NewDecisionRequest()
	request.Placements = []model.Placement{placement}
	request.Keywords = []string{"keyword1", "keyword2"}
	request.User = model.NewUser("abc")

	decisions := client.Decisions()
	response, _ := decisions.Get(request)

	s, _ := json.MarshalIndent(response, "", "  ")
	fmt.Println(string(s))
}
```

### Recording Impression & Clicks

Use with the fetch ad example above.

```go
decision := response.Decisions["div0"]

// Impression pixel; fire when user sees the ad
pixels := client.Pixels()
pixels.Fire(NewPixelFireOptions(decision.ImpressionUrl))

// Click pixel; fire when user clicks on the ad
// status: HTTP status code
// location: click target URL
clickResponse, _ := pixels.Fire(NewPixelFireOptions(decision.ClickUrl))
fmt.Printf("Fired! status: %d location: %s\n", clickResponse.StatusCode, clickResponse.Location)
```

### UserDB: Reading User Record

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/linksports/kevel-decision-sdk-go"
)

func main() {
	// Demo network ID; find your own via the Kevel UI!
	opts := kevel.NewClientOptions(23)
	client := kevel.NewClient(opts)
	userDb := client.UserDb()
	record, _ := userDb.Read(userKey)

	s, _ := json.MarshalIndent(record, "", "  ")
	fmt.Println(string(s))
}
```

### UserDB: Setting Custom Properties

```go
package main

import "github.com/linksports/kevel-decision-sdk-go"

func main() {
	// Demo network ID; find your own via the Kevel UI!
	opts := kevel.NewClientOptions(23)
	client := kevel.NewClient(opts)
	userDb := client.UserDb()

	props := map[string]interface{}{
		"favoriteColor":  "blue",
		"favoriteNumber": 42,
		"favoriteFoods":  []string{"strawberries", "chocolate"},
	}

	userDb.SetCustomProperties("abc", props)
}
```

### UserDB: Forgetting User Record

```go
package main

import "github.com/linksports/kevel-decision-sdk-go"

func main() {
	// Demo network ID and API key; find your own via the Kevel UI!
	opts := kevel.NewClientOptions(23)
	opts.ApiKey = "YOUR-API-KEY"
	client := kevel.NewClient(opts)
	userDb := client.UserDb()
	userDb.Forget("abc")
}
```

## Documentation

- [Kevel API Documentation](https://dev.kevel.co/reference)
- [Kevel User & Developer Documentation](https://dev.kevel.co/docs)

## Contributing

### Reporting Issues

- For bug fixes and improvements to this SDK please use Github to [open an issue](https://github.com/linksports/kevel-decision-sdk-go/issues) or send us a [pull request](https://github.com/linksports/kevel-decision-sdk-go/pulls).

### Building / Running Tests

```
go build
go test
```
