package endpoints

func SetEndpoint(endpointName string, endpoint *Endpoint, mdw map[string][]Middleware) {
	for _, m := range mdw[endpointName] {
		*endpoint = m(*endpoint)
	}
}
