package main

import "net/http"

func handlerJSON(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
