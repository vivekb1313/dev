package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	URLs []string `json:"urls"`
}

type Response struct {
	Status map[string]string `json:"status"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	results := checkURLs(req.URLs)

	res := Response{Status: results}
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(res)

}
