package middleware

import "net/http"

type CostMiddleware struct {
}

func NewCostMiddleware() *CostMiddleware {
	return &CostMiddleware{}
}

func (m *CostMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		// Passthrough to next handler if need
		next(w, r)
	}
}
