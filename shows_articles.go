package gotaseries

type ShowsArticlesOptions struct {
	ID int `url:"id"`
}

func (s *ShowService) Articles(opts ShowsArticlesOptions) ([]Article, error) {
	u, err := addOptions("/shows/articles", opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var articles *ArticlesResponse
	_, err = s.client.do(req, &articles)

	if err = articles.Errors.Err(); err != nil {
		return nil, err
	}

	return articles.Articles, nil
}