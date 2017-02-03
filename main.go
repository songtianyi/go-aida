package main

import (
	"github.com/songtianyi/wechat-go"
)

func main() {
	wxbot.AutoLogin()
	wxbot.Run()
}
