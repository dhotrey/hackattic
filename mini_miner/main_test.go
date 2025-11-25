package main

import (
	"encoding/hex"
	"testing"

	"github.com/charmbracelet/log"
)

func TestSolvedfunc(t *testing.T) {
	tests := map[int]string{
		8:  "00d696db487caf06a2f2a8099479577c3785c37b3d8a77dc413cfb19ec2e0141",
		16: "000096db487caf06a2f2a8099479577c3785c37b3d8a77dc413cfb19ec2e0141",
		13: "000696db487caf06a2f2a8099479577c3785c37b3d8a77dc413cfb19ec2e0141",
	}

	for exp, h := range tests {
		b, err := hex.DecodeString(h)
		if err != nil {
			log.Error("Invalid hex")
			log.Error("", "exp", exp, "h", h)
			log.Fatal(err)
		}

		got := solved([32]byte(b), exp, getNewLogger("test", log.DebugLevel))
		expected := true
		if got != expected {
			t.Errorf("expected %v , got %v for key : %d", expected, got, exp)
		}
	}
}
