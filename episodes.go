package gotaseries

type EpisodesResponse struct {
	Episodes []Episode `json:"episodes"`
	Errors   `json:"errors"`
}
type Note struct {
	Total   int     `json:"total"`
	Mean    float64 `json:"mean"`
	User    int     `json:"user"`
	Moyenne float64 `json:"moyenne"`
}

type WatchedBy struct {
	ID    int      `json:"id"`
	Login string   `json:"login"`
	Note  *float64 `json:"note"`
}

type Episode struct {
	ID          int      `json:"id"`
	ThetvdbID   int      `json:"thetvdb_id"`
	YoutubeID   *string  `json:"youtube_id"`
	Title       string   `json:"title"`
	Season      int      `json:"season"`
	Episode     int      `json:"episode"`
	Code        string   `json:"code"`
	Global      int      `json:"global"`
	Description string   `json:"description"`
	Director    string   `json:"director"`
	Writers     []string `json:"writers"`
	Special     int      `json:"special"`
	Comments    string   `json:"comments"`
	ResourceURL string   `json:"resource_url"`
	Note        Note     `json:"note"`
	User        struct {
		Seen       bool `json:"seen"`
		Hidden     bool `json:"hidden"`
		Downloaded bool `json:"downloaded"`
	} `json:"user"`
	Date         string        `json:"date"`
	WatchedBy    []WatchedBy   `json:"watched_by"`
	SeenTotal    int           `json:"seen_total"`
	ReleasesSvod []interface{} `json:"releasesSvod"`
	Show         struct {
		ID        int    `json:"id"`
		ThetvdbID int    `json:"thetvdb_id"`
		Title     string `json:"title"`
		InAccount bool   `json:"in_account"`
	} `json:"show"`
	Subtitles *[]Subtitle `json:"subtitles"`
}
