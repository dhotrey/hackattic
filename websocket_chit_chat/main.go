package main

import (
	"fmt"
	"io"
	"math"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	log := getNewLogger("main")
	log.Info("WebSocket chit chat!")
	tok := getChalToken()
	log.Info("Got challenge token", "token", tok)

	u := url.URL{
		Scheme: "wss",
		Host:   "hackattic.com",
		Path:   fmt.Sprintf("/_/ws/%s", tok),
	}

	log.Info("challenge endpoint", "url", u.String())
	c, r, err := websocket.DefaultDialer.Dial(u.String(), nil)
	startTime := time.Now()
	if err != nil {
		log.Error(r.Status)
		body, _ := io.ReadAll(r.Body)
		log.Error(string(body))
		log.Fatal(err)
	}
	defer c.Close()

	for i := 0; ; i++ {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		m := string(msg)
		log.Info(m)
		timeElapsed := time.Since(startTime).Abs().Milliseconds()
		if m == "ping!" {
			closestDuration, wsMsg := GetClosestDuration(timeElapsed)
			log.Infof("Time elapsed %d , closest duration %d", timeElapsed, closestDuration)
			c.WriteMessage(websocket.TextMessage, wsMsg)
			startTime = time.Now()
		}
	}
}

func GetClosestDuration(t int64) (int64, []byte) {
	intervals := []int64{700, 1500, 2000, 2500, 3000}
	closestIdx := 0
	minDiff := math.MaxFloat32

	for i, interval := range intervals {
		diff := math.Abs(float64(t - interval))
		if diff < minDiff {
			minDiff = diff
			closestIdx = i
		}
	}
	val := intervals[closestIdx]
	msgString := strconv.FormatInt(val, 10)
	return val, []byte(msgString)
}
