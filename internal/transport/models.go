package transport

type CategoryName = string

type CategoryViews struct {
	CategoryName CategoryName `json:"categoryName"`
	Views        int          `json:"views"`
}

type GetPieChartStatsResponse struct {
	AdvertisingId         string          `json:"advertisingId"`
	TotalViewsPerCategory []CategoryViews `json:"totalViewsPerCategory"`
}

type DailyView struct {
	AdvertisingId string               `json:"advertisingId"`
	Date          string               `json:"date"`
	Views         map[CategoryName]int `json:"views"`
}
