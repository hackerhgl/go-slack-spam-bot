package main

import (
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.ActiveName("Slack")
	count := 25
	for i := 0; i < count; i++ {
		str := "GOLANG BOT. index:" + strconv.Itoa(i) + "; time:" + time.Now().String()
		robotgo.TypeStr(str)
		robotgo.KeyTap("enter")
		time.Sleep(time.Millisecond * 200)
	}

}
