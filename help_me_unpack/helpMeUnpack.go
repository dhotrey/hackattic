package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hackattic/utils"

	"math"

	l "github.com/charmbracelet/log"
)

const chalName = "help_me_unpack"

var log *l.Logger

type chalBody struct {
	B64Str string `json:"bytes"`
}

type Sol struct {
	SignedInt          int32   `json:"int"`
	UnsignedInt        uint32  `json:"uint"`
	SignedShortVal     int16   `json:"short"`
	FloatVal           float32 `json:"float"`
	DoubleVal          float64 `json:"double"`
	DoubleBigEndianVal float64 `json:"big_endian_double"`
}

func init() {
	log = utils.GetLogger(chalName)
}

func main() {
	b := chalBody{}
	s := Sol{}
	resp := utils.GetChal(chalName)
	log.Debug(string(resp))
	json.Unmarshal(resp, &b)
	log.Debug(b)
	byteArray, err := base64.StdEncoding.DecodeString(b.B64Str)
	if err != nil {
		log.Errorf("err: %v\n", err)
	}
	log.Debugf("rawDecodedTxt: %v\n", byteArray)
	log.Infof("len(byteArray):%d", len(byteArray))
	s.SignedInt = int32(binary.LittleEndian.Uint32(byteArray[0:4]))
	s.UnsignedInt = binary.LittleEndian.Uint32(byteArray[4:8])
	s.SignedShortVal = int16(binary.LittleEndian.Uint16(byteArray[8:12]))
	s.FloatVal = math.Float32frombits(binary.LittleEndian.Uint32(byteArray[12:16]))
	s.DoubleVal = math.Float64frombits(binary.LittleEndian.Uint64(byteArray[16:24]))
	s.DoubleBigEndianVal = math.Float64frombits(binary.BigEndian.Uint64(byteArray[24:32]))
	jsonData, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	log.Infof("s -> %v", s)
	log.Debug(string(jsonData))
	status := utils.SendSol(chalName, jsonData)
	fmt.Printf("status: %v\n", status)
}
