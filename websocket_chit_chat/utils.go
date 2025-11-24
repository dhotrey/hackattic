package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func getChalToken() string {
	resp, err := http.Get("https://hackattic.com/challenges/websocket_chit_chat/problem?access_token=9e115aa83183d27a")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var response struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal(err)
	}
	return response.Token
}

func getNewLogger(prefix string) *log.Logger {
	return log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      "01:04:05.000",
		Level:           log.DebugLevel,
		ReportTimestamp: true,
		Prefix:          prefix,
	})
}
