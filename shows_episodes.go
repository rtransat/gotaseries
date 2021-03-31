package gotaseries

type ShowsEpisodesOptions struct {
	ID        int `url:"id,omitempty"`
	ThetvdbID int `url:"thetvdb_id,omitempty"`
	Season    int `url:"season,omitempty"`
	Episode   int `url:"episode,omitempty"`
	Subtitles bool `url:"subtitles,omitempty"`
}

func (s *ShowService) Episodes(opts ShowsEpisodesOptions) ([]Episode, error) {
	u, err := addOptions("/shows/episodes", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var episodes EpisodesResponse
	_, err = s.client.do(req, &episodes)
	if err != nil {
		return nil, err
	}

	if err = episodes.Errors.Err(); err != nil {
		return nil, err
	}

	return episodes.Episodes, nil
}
