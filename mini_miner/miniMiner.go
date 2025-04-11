package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"hackattic/utils"
	"sort"
	"strings"

	l "github.com/charmbracelet/log"
)

const chalName = "mini_miner"

var log *l.Logger

type chalBody struct {
	Difficulty int   `json:"difficulty"`
	Block      Block `json:"block"`
}

type Block struct {
	Nonce string          `json:"nonce"`
	Data  [][]interface{} `json:"data"`
}

type Sol struct {
	Nonce string `json:"nonce"`
}

func init() {
	log = utils.GetLogger(chalName)
}

func main() {
	h := sha256.New()
	b := chalBody{}
	// s := Sol{}
	resp := utils.GetChal(chalName)
	// log.Debug(string(resp))
	json.Unmarshal(resp, &b)
	// log.Debug(b.Block.Data)
	sort.Slice(b.Block.Data, func(i, j int) bool { // sort data alphabetically
		valI := b.Block.Data[i][0].(string)
		valJ := b.Block.Data[j][0].(string)
		return valI < valJ
	})
	// log.Debug(b.Block.Data)
	blockDataJSON, err := json.Marshal(b.Block.Data)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10000; i++ {
		hashable := fmt.Sprintf("{\"data\":%s,\"nonce\":%d}", blockDataJSON, i)
		h.Write([]byte(hashable))
		sum := h.Sum(nil)
		hexHash := fmt.Sprintf("%x", sum)
		binHash := fmt.Sprintf("%b", sum)
		fmt.Println("-----------------------------------")
		fmt.Printf("difficulty : %d\n", b.Difficulty)
		fmt.Printf("Sha256 : %s\n", hexHash)
		fmt.Printf("Sha256 (binary) : %s\n", binHash)
		fmt.Printf("nounce : %d\n", i)
		valid := checkLeadingBits(binHash, b.Difficulty)
		fmt.Println("-----------------------------------")
		fmt.Printf("\n\n")
		if valid {
			return
		}
	}
}

func checkLeadingBits(hash string, d int) bool {
	bitsStr := hash[1 : d+2] // compensating for 2 spaces
	splitBits := strings.Split(bitsStr, " ")
	bits := strings.Join(splitBits, "")
	fmt.Printf("bitsStr: (%d) %v\n", len(bits), bits)
	for _, b := range bits {
		if b != 48 {
			return false
		}
	}
	return true
}
