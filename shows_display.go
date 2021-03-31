package gotaseries

type ShowsDisplayOptions struct {
	ID        int    `url:"id,omitempty"`
	ThetvdbID int    `url:"thetvdb_id,omitempty"`
	ImdbID    int    `url:"imdb_id,omitempty"`
}

func (s *ShowService) Display(opts ShowsDisplayOptions) (*Show, error) {
	u, err := addOptions("/shows/display", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var show ShowResponse
	_, err = s.client.do(req, &show)
	if err != nil {
		return nil, err
	}

	if err = show.Errors.Err(); err != nil {
		return nil, err
	}

	return &show.Show, nil
}