package dtos

type DashboardResponse struct {
	DashboardStatus DashboardStats `json:"dashboard_status"`
	PeakHour        []PeakHour     `json:"peak_hour"`
}

type DashboardStats struct {
	TotalUniqueVisitors int64 `json:"total_unique_visitors"`
	TotalActiveToday    int64 `json:"total_active_today"`
	TotalVisitorsToday  int64 `json:"total_visitors_today"`
}

type PeakHour struct {
	Hour        int   `json:"hour"`
	TotalVisits int64 `json:"total_visits"`
}
