package schedule

type ScheduleData struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	PublishedAt string `json:"published_at"`
	AvailableAt string `json:"available_at"`
	Duration number `json:"duration"`
	Status string `json:"status"`
}