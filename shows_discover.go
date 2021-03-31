package gotaseries

type ShowsDiscoverOptions struct {
	Limit   int  `url:"limit,omitempty"`
	Offset  int  `url:"offset,omitempty"`
	Summary bool `url:"summary,omitempty"`
}

func (s *ShowService) Discover(opts *ShowsDiscoverOptions) ([]Show, error) {
	u, err := addOptions("/shows/discover", opts)
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