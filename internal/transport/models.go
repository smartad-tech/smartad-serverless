package transport

type CategoryViews struct {
	CategoryName string `json:"categoryName"`
	Views        int    `json:"views"`
}

type GetAdStatsResponse struct {
	AdvertisingId         string          `json:"advertisingId"`
	TotalViewsPerCategory []CategoryViews `json:"totalViewsPerCategory"`
}
