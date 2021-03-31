package gotaseries

type SubtitlesResponse struct {
	Subtitles []Subtitle `json:"subtitles"`
	Errors    `json:"errors"`
}

type Subtitle struct {
	ID       int           `json:"id"`
	Language string        `json:"language"`
	Source   string        `json:"source"`
	Quality  int           `json:"quality"`
	File     string        `json:"file"`
	Content  []interface{} `json:"content"`
	URL      string        `json:"url"`
	Episode  *struct {
		ShowID    int `json:"show_id"`
		EpisodeID int `json:"episode_id"`
		Season    int `json:"season"`
		Episode   int `json:"episode"`
	} `json:"episode"`
	Date string `json:"date"`
}
