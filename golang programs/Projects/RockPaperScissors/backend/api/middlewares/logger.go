package middlewares

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("%v %v %v", r.Method, r.URL, r.Host)
		next.ServeHTTP(w, r)
	})
}
