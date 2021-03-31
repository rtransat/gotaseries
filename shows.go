package gotaseries

type ShowsResponse struct {
	Shows  []Show `json:"shows"`
	Errors `json:"errors"`
}
type ShowResponse struct {
	Show   Show `json:"show"`
	Errors `json:"errors"`
}

type SeasonsDetail struct {
	Number   int `json:"number"`
	Episodes int `json:"episodes"`
}

type Showrunner struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type Notes struct {
	Total int     `json:"total"`
	Mean  float64 `json:"mean"`
	User  int     `json:"user"`
}

type Images struct {
	Show   string `json:"show"`
	Banner string `json:"banner"`
	Box    string `json:"box"`
	Poster string `json:"poster"`
}

type Next struct {
	ID    interface{} `json:"id"`
	Code  string      `json:"code"`
	Date  interface{} `json:"date"`
	Title interface{} `json:"title"`
	Image interface{} `json:"image"`
}

type FriendsWatching struct {
	ID     int         `json:"id"`
	Login  string      `json:"login"`
	Note   interface{} `json:"note"`
	Avatar string      `json:"avatar"`
}

type User struct {
	Archived        bool              `json:"archived"`
	Favorited       bool              `json:"favorited"`
	Remaining       int               `json:"remaining"`
	Status          int               `json:"status"`
	Last            string            `json:"last"`
	Tags            interface{}       `json:"tags"`
	Next            Next              `json:"next"`
	FriendsWatching *[]FriendsWatching `json:"friends_watching"`
}

type Available struct {
	Last int `json:"last"`
}

type Svod struct {
	ID        interface{} `json:"id"` // https://www.betaseries.com/bugs/api/453
	Name      string      `json:"name"`
	Tag       string      `json:"tag"`
	LinkURL   string      `json:"link_url"`
	Available Available   `json:"available"`
	Logo      string      `json:"logo"`
	Partner   bool        `json:"partner"`
}

type Platforms struct {
	Svods []Svod `json:"svods"`
}

type Show struct {
	ID             int                `json:"id"`
	ThetvdbID      int                `json:"thetvdb_id"`
	ImdbID         string             `json:"imdb_id"`
	Title          string             `json:"title"`
	OriginalTitle  *string            `json:"original_title"`
	Description    *string            `json:"description"`
	Seasons        string             `json:"seasons"`
	SeasonsDetails *[]SeasonsDetail   `json:"seasons_details"`
	Episodes       string             `json:"episodes"`
	Followers      string             `json:"followers"`
	Comments       *string            `json:"comments"`
	Similars       *string            `json:"similars"`
	Characters     *string            `json:"characters"`
	Creation       string             `json:"creation"`
	Showrunner     *Showrunner        `json:"showrunner"`
	Genres         *map[string]string `json:"genres"`
	Length         *string            `json:"length"`
	Network        *string            `json:"network"`
	Country        *interface{}       `json:"country"`
	Rating         *string            `json:"rating"`
	Status         *string            `json:"status"`
	Language       *string            `json:"language"`
	Notes          *Notes             `json:"notes"`
	InAccount      *bool              `json:"in_account"`
	Images         *Images            `json:"images"`
	Aliases        *map[int]string    `json:"aliases"`
	SocialLinks    *[]interface{}     `json:"social_links"`
	User           *User              `json:"user"`
	NextTrailer    *string            `json:"next_trailer"`
	ResourceURL    *string            `json:"resource_url"`
	Platforms      *Platforms          `json:"platforms"`
}

type ShowService service