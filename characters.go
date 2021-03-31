package gotaseries

type CharactersResponse struct {
	Characters []Character `json:"characters"`
	Errors     `json:"errors"`
}
type Character struct {
	ID          int     `json:"id"`
	ShowID      int     `json:"show_id"`
	Name        string  `json:"name"`
	Role        string  `json:"role"`
	Actor       string  `json:"actor"`
	Picture     *string `json:"picture"`
	Description *string `json:"description"`
}
