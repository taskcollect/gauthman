package handlers

import (
	"fmt"
	"log"
	"main/auth"
	"net/http"
	"strconv"

	"github.com/buger/jsonparser"
)

func (h *BaseHandler) ExchangeToken(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if !q.Has("code") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing 'code' query parameter"))
		return
	}

	code := q.Get("code")
	token, err := auth.ExchangeCodeForTokenPair(code, h.Secrets)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	resp := []byte{'{', '}'}

	resp, err = jsonparser.Set(resp, []byte(strconv.Quote(token.AccessToken)), "access")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	resp, err = jsonparser.Set(resp, []byte(strconv.Quote(token.RefreshToken)), "refresh")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	resp, err = jsonparser.Set(resp, []byte(fmt.Sprint(token.Expiry.UTC().Unix())), "expires")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
