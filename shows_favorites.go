package gotaseries

type ShowsFavoritesOptions struct {
	ID      int    `url:"id,omitempty"`
	Order   string `url:"order,omitempty"`
	Limit   int    `url:"limit,omitempty"`
	Offset  int    `url:"offset,omitempty"`
	Status  string `url:"offset,omitempty"`
	Summary bool   `url:"summary,omitempty"`
}

func (s *ShowService) Favorites(opts *ShowsFavoritesOptions) ([]Show, error) {
	u, err := addOptions("/shows/favorites", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var shows *ShowsResponse
	_, err = s.client.do(req, &shows)

	if err = shows.Errors.Err(); err != nil {
		return nil, err
	}

	return shows.Shows, nil
}
