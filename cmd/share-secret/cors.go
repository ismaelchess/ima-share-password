package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	// DefaultCORSMaxAge is the recommended TTL for a CORS cache
	DefaultCORSMaxAge = time.Hour * 24

	// DefaultCORSMethods are the minimal methods CORS needs to support
	DefaultCORSMethods = []string{
		http.MethodGet,
	}

	// DefaultCORSHeaders allows all headers for CORS middleware
	DefaultCORSHeaders = []string{"*"}
)

func CORS(maxAge time.Duration, headers []string, methods []string) Middleware {
	allowedHeaders := strings.Join(headers, ",")
	allowedMethods := strings.Join(methods, ",")
	maxAgeInSeconds := strconv.FormatInt(int64(maxAge/time.Second), 10)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
			w.Header().Set("Access-Control-Max-Age", maxAgeInSeconds)
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func CORSWithDefaults() Middleware {
	return CORS(DefaultCORSMaxAge, DefaultCORSHeaders, DefaultCORSMethods)
}
