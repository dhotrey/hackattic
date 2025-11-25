package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io"
	"math/bits"
	"net/http"

	"github.com/charmbracelet/log"
)

func main() {
	log := getNewLogger("main", log.InfoLevel)
	log.Info("Mini Miner!")
	resp, err := http.Get("https://hackattic.com/challenges/mini_miner/problem?access_token=9e115aa83183d27a")
	if err != nil {
		log.Fatal(err)
	}
	problemBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(string(problemBody))

	var chal map[string]any
	err = json.Unmarshal(problemBody, &chal)

	difficulty := chal["difficulty"].(float64)
	log.Infof("Got difficulty : %v", difficulty)

	block := chal["block"]
	j, err := json.Marshal(block)
	log.Infof("Block %s", j)
	sha := sha256.Sum256(j)
	log.Info(sha)
	log.Infof("Initial Hash - nonce : null | hash : %x", sha)
	log.Infof("Starting mining for difficulty %v", difficulty)
	nonce := 0

	for nonce = 0; !solved(sha, int(difficulty), log); nonce++ {
		rawBlock := chal["block"]
		if blockMap, ok := rawBlock.(map[string]any); ok {
			blockMap["nonce"] = nonce
			jason, err := json.Marshal(blockMap)
			if err != nil {
				log.Fatal(err)
			}
			sha = sha256.Sum256(jason)
			log.Infof("nonce : %d | hash : %x", nonce, sha)
		}
	}

	log.Infof("Solved with nonce %d", nonce)

	solutionBody := map[string]int{
		"nonce": nonce-1,
	}

	body, err := json.Marshal(solutionBody)
	if err != nil {
		log.Error(err)
	}
	bodyReader := bytes.NewBuffer(body)

	resp, err = http.Post("https://hackattic.com/challenges/mini_miner/solve?access_token=9e115aa83183d27a", "application/json", bodyReader)
	if err != nil {
		log.Error(err)
	}
	postResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(string(postResp))
}

func solved(hash [32]byte, difficulty int, l *log.Logger) bool {
	leadingZeros := 0
	l.Debugf("Difficulty %d", difficulty)
	for i, byt := range hash {
		leadingZeros += bits.LeadingZeros8(uint8(byt))
		l.Debug("", "i+1", (i + 1), "byt", byt, "leadingZeros", leadingZeros)
		if leadingZeros == difficulty {
			return true
		} else {
			if (i+1)*8 > difficulty { // when number of bits encountered is greater than difficulty
				l.Debug("number of bits encountered is greater than difficulty")
				return false //  but leadingZeros is not equal to difficulty
			} else if leadingZeros != (i+1)*8 { // when number of bits encountered is less than difficulty
				l.Debug("number of bits encountered is less than difficulty")
				return false // but not all bits encountered are zero
			}
		} // continue iteration
		l.Debug("continuing iteration")
	}
	return false
}
