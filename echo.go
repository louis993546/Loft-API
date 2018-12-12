package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// TimeResponse contains time in string + format that the time string is in.
type TimeResponse struct {
	Time   string `json:"time"`
	Format string `json:"format"`
}

// AllEcho handles GET/PUT/POST/PATCH/DELETE echo/ by returning the server time
// It's useful for anyone to check if the server is online
func AllEcho(w http.ResponseWriter, r *http.Request) {
	timeNow := TimeResponse{time.Now().UTC().Format(time.RFC3339), "RFC3339"}
	json.NewEncoder(w).Encode(timeNow)
}
