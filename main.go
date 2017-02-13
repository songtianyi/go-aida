package main

import (
	"github.com/songtianyi/wechat-go"
	"github.com/songtianyi/rrframework/connector/redis"
	"time"
	"strings"
	"fmt"
	"math/rand"
)

func delText(msg map[string]interface{}) {
	content := msg["Content"].(string)
	fmt.Println(content)
	FromUserName := msg["FromUserName"].(string)
	//if strings.Contains(FromUserName, "@@") {
	//	// ignore group message
	//	return
	//}
	_, rc := rrredis.GetRedisClient("10.19.147.75:6379")
	//aszone, _ := time.LoadLocaltion("Asia/Shanghai")
	now := time.Now().Unix()
	if strings.Contains(strings.ToLower(content), "sleep") {
		rc.ZAddBatch("SLEEPING:RECORD:songtianyi:NIGHT", []float64{float64(now)}, []interface{}{float64(now)})
		wxbot.SendText("good night :)", wxbot.Bot.UserName, FromUserName)
	}
	if strings.Contains(strings.ToLower(content), "awake") {
		rc.ZAddBatch("SLEEPING:RECORD:songtianyi:MORNING", []float64{float64(now)}, []interface{}{float64(now)})
		wxbot.SendText("good morning :)", wxbot.Bot.UserName, FromUserName)
	}
	if strings.Contains(strings.ToLower(content), "record") {
		resm, err := rc.ZRevRange("SLEEPING:RECORD:songtianyi:MORNING", 0, 31)
		if err != nil {
			fmt.Println(err)
			wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
		}
		fmt.Println(resm)
		resn, err := rc.ZRevRange("SLEEPING:RECORD:songtianyi:NIGHT", 0, 31)
		if err != nil {
			fmt.Println(err)
			wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
		}
		fmt.Println(resn)
	}
	var paths = []string{
		"/data/sexx/haixiuzu/p10165173.jpg",
		"/data/sexx/haixiuzu/p10165173.jpg",
		"/data/sexx/haixiuzu/p67467920.jpg",
		"/data/sexx/haixiuzu/p67548323.jpg",
		"/data/sexx/haixiuzu/p67616723.jpg",
		"/data/sexx/haixiuzu/p67058029.jpg",
		"/data/sexx/haixiuzu/p67511086.jpg",
		"/data/sexx/haixiuzu/p67555235.jpg",
		"/data/sexx/haixiuzu/p67616725.jpg",
		"/data/sexx/haixiuzu/p67069563.jpg",
		"/data/sexx/haixiuzu/p67533993.jpg",
		"/data/sexx/haixiuzu/p67616714.jpg",
		"/data/sexx/haixiuzu/p67616732.jpg",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if strings.Contains(strings.ToLower(content), "fache") {
		wxbot.SendImg(paths[r.Intn(len(paths))], wxbot.Bot.UserName, FromUserName)
	}
}

func main() {
	// add text message handler
	wxbot.HandlerRegister.Add(1, wxbot.Handler(delText))
	// login
	wxbot.AutoLogin()
	// enter message loop
	wxbot.Run()
}
