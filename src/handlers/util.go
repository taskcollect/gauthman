package handlers

import (
	"log"
	"main/auth"
	"net/http"
	"time"
)

// handler for the base of the api that allows it to access the db object
type BaseHandler struct {
	Secrets *auth.OAuth2Secrets
}

// function to make a new base handler given a db
func NewBaseHandler(secrets *auth.OAuth2Secrets) *BaseHandler {
	return &BaseHandler{
		Secrets: secrets,
	}
}

// make sure that the method of the request is what we expect
func EnsureMethod(method string, w http.ResponseWriter, r *http.Request) bool {
	// this must be a post request
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, expected " + method))
		return false
	}
	return true
}

func RequestLogger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		mux.ServeHTTP(w, r)

		log.Printf(
			"%s %s from [ %s ] done in [ %v ]",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
