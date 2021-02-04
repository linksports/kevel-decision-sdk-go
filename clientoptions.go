package kevel

type ClientOptions struct {
	NetworkId int
	SiteId    int
	Protocol  string
	Host      string
	Path      string
	ApiKey    string
}

func NewClientOptions(networkId int) ClientOptions {
	return ClientOptions{NetworkId: networkId}
}
