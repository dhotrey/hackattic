package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func GetChal(chalName string) []byte {
	chalUrl := fmt.Sprintf("https://hackattic.com/challenges/%s/problem?access_token=9e115aa83183d27a", chalName)
	resp, err := http.Get(chalUrl)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return body
}

func SendSol(chalName string, jsonData []byte) string {
	chalUrl := fmt.Sprintf("https://hackattic.com/challenges/%s/solve?access_token=9e115aa83183d27a", chalName)
	resp, err := http.Post(chalUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
    fmt.Printf("body: %s\n", string(body))

	return resp.Status
}
