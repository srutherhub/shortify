package models

import "net/http"

type ServerError string

type Route struct {
	Path     string
	Method   string
	Handler  http.HandlerFunc
	IsPublic bool
}

type ServerErrorResponse struct {
	Error   ServerError `json:"error"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

type ShortifyRequestBody struct {
	URL         string `json:"url"`
	ExpiresIn   int64  `json:"expires_in"`
	UtmSource   string `json:"utm_source"`
	UtmMedium   string `json:"utm_medium"`
	UtmCampaign string `json:"utm_campaign"`
	UtmTerm     string `json:"utm_term"`
	UtmContent  string `json:"utm_content"`
}

type ShortifyResponse struct {
	URL string `json:"shorturl"`
}

type QuickStatsResponse struct {
	Total_Urls             int64   `json:"total_url_count"`
	Total_Click_Count      int64   `json:"total_click_count"`
	Average_Clicks_Per_Url float64 `json:"avg_clicks_per_url"`
}

const (
	EInvalidRequest ServerError = "INVALID_REQUEST"
	EAppError       ServerError = "APP_ERROR"
	EDatabaseError  ServerError = "DATABASE_ERROR"
)
