package gotaseries

type ArticlesResponse struct {
	Articles []Article `json:"articles"`
	Locale   string    `json:"locale"`
	Errors   `json:"errors"`
}
type Article struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Excerpt *string `json:"excerpt"`
	Content string `json:"content"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Image   string `json:"image"`
	Sticky  string `json:"sticky"`
}
