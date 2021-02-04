package kevel

type PixelFireOptions struct {
	Url               string
	RevenueOverride   *float64
	AdditionalRevenue *float64
	EventMultiplier   *int
}

func NewPixelFireOptions(url string) PixelFireOptions {
	return PixelFireOptions{Url: url}
}
