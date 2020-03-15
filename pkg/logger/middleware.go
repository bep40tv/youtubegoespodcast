package logger

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// Middleware act as middleware and log request method and path
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/metrics" {
			logrus.WithField("headers", r.Header).Infof("[API] Request %s %s", r.Method, r.RequestURI)
		}

		next.ServeHTTP(w, r)
	})
}
