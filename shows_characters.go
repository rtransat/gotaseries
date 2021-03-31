package gotaseries

type ShowsCharactersOptions struct {
	ID        int `url:"id,omitempty"`
	ThetvdbID int `url:"thetvdb_id,omitempty"`
}

func (s *ShowService) Characters(opts *ShowsCharactersOptions) ([]Character, error) {
	u, err := addOptions("/shows/characters", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var characters *CharactersResponse
	_, err = s.client.do(req, &characters)

	if err = characters.Errors.Err(); err != nil {
		return nil, err
	}

	return characters.Characters, nil
}
