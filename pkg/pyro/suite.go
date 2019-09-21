package pyro

type Route struct {
	Url    string
	Status int
}

type Suite struct {
	Routes []Route
}

func NewSuite(urls ...string) Suite {
	s := Suite{}
	for _, u := range urls {
		s.AddRoute(Route{
			Url:    u,
			Status: 200,
		})
	}

	return s
}

func (s *Suite) AddRoute(r Route) {
	s.Routes = append(s.Routes, r)
}
