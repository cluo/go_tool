package mylog

import (
	"testing"
)

func TestLog(t *testing.T) {
	// log filename
	logger:=NewLog("test.log")

	// log level
	logger.SetLevel("debug")

	logger.Println("1.println","11.println")
	logger.Error("error")
	logger.Debugf("%v","debug")
}
