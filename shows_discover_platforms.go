package gotaseries

type ShowsDiscoverPlatformsOptions struct {
	Summary bool `url:"summary,omitempty"`
}

func (s *ShowService) DiscoverPlatforms(opts *ShowsDiscoverPlatformsOptions) ([]Show, error) {
	u, err := addOptions("/shows/discover_platforms", opts)
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