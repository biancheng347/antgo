package test

import (
	"github.com/small-ek/antgo/aemail"
	"testing"
)

func TestEmail(t *testing.T) {
	aemail.New("56494565@qq.com").SetTo([]string{"56494565@qq.com"}).SetTitle("test").SetText("test2223223").SetPassword("fdtshicbbvybbiic").SetFilePath([]string{"test.txt"}).Send()
}
