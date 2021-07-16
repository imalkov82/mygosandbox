package main

import (
	"encoding/json"
	"net/http"
)

type JobPost struct {
	Vendor string `json:"vendor"`
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	j := JobPost{}
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(j.Vendor))
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8000", nil)
}
