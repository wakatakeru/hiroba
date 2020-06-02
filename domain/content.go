package domain

type Content struct {
	ID     int    `json:"id"`
	SiteID int    `json:"site_id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
