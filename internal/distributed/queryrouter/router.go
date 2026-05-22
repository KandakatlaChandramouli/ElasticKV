package queryrouter

type Route struct {
	Node  string
	Query string
}

func RouteQuery(
	query string,
	nodes []string,
) []Route {

	routes := make([]Route, 0)

	for _, node := range nodes {

		routes = append(
			routes,
			Route{
				Node:  node,
				Query: query,
			},
		)
	}

	return routes
}
