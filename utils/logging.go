package utils

import (
	"os"

	"github.com/charmbracelet/log"
)

var chalName string

func init() {
	// l := log.New(os.Stderr)
	// l.SetLevel(log.DebugLevel)
	// fname := fmt.Sprintf("%v.log", chalName)
	// file, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	l.Fatal(err)
	// }
	// os.Stderr = file
	// fmt.Fprintln(os.Stderr, "-------------------------------------------------")
}

func GetLogger(c string) *log.Logger {
	chalName = c
	l := log.New(os.Stdout)
	l.SetLevel(getLogLevel())
	// fname := fmt.Sprintf("%v.log", chalName)
	// file, err := os.OpenFile(fname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	// if err != nil {
	// 	l.Fatal(err)
	// }
	//    os.Stderr = file
	return l
}

func getLogLevel() log.Level {
	if len(os.Args) < 2 {
		return log.DebugLevel
	} else {
		ll := os.Args[1]
		if ll == "info" || ll == "i" || ll == "inf" {
			return log.InfoLevel
		}
	}
	return log.DebugLevel
}
