package middleware

import "net/http"

type Middleware struct {
	httpHandler http.Handler
}

func NewMiddlware(httpHandler http.Handler) *Middleware {
	return &Middleware{httpHandler: httpHandler}
}

func (s *Middleware) CorsMiddleware() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		s.httpHandler.ServeHTTP(w, r)
	})
}
