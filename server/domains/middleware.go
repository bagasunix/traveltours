package domains

type Middleware func(contract Service) Service
