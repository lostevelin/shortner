package main

import (
	"io"
	"net/http"
)

var list = map[string]string{}

func postHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		return
	}
	url := "http://localhos:8080/EwHXdJfB"

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	list["EwHXdJfB"] = string(body)

	res.Header().Set("content-type", "text/plain")
	res.Header().Set("content-length", "30")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(url))
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		return
	}
	http.Redirect(res, req, list["EwHXdJfB"], http.StatusTemporaryRedirect)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", postHandler)
	mux.HandleFunc("/EwHXdJfB", getHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
