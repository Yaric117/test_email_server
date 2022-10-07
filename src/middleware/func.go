package middleware

import (
	"context"
	"net/http"
	"strings"
	"testproject/config"
	apiV1 "testproject/controller/http/v1"
)

var AppTokenAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			next.ServeHTTP(w, r)
		}

		var payload interface{}

		apiUrl := strings.HasPrefix(r.URL.Path, "/api")

		if !apiUrl {
			next.ServeHTTP(w, r)
			return
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			apiV1.RespondError(w, "Auth token is empty!", payload, http.StatusUnauthorized)
			return
		}

		if tokenHeader != config.App.Token {
			apiV1.RespondError(w, "Not correct authentication token!", payload, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", "client")
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

var AddCors = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowMethods := []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions}

		w.Header().Add("Access-Control-Allow-Origin", config.App.Cors.AllowedOrigins)
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", strings.Join(allowMethods, ","))

		next.ServeHTTP(w, r)
	})
}
