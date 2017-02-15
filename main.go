package main

import (
	"github.com/songtianyi/wechat-go"
	"github.com/songtianyi/rrframework/connector/redis"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/rrframework/storage"
	"github.com/songtianyi/laosj/spider"
	"time"
	"strings"
	"fmt"
	"math/rand"
	"net/url"
	"net/http"
	"io/ioutil"
)

func delText(msg map[string]interface{}) {
	content := msg["Content"].(string)
	fmt.Println(content)
	FromUserName := msg["FromUserName"].(string)
	ToUserName := msg["ToUserName"].(string)

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
		target := FromUserName
		if wxbot.Bot.UserName == FromUserName {
			target = ToUserName
		}
		wxbot.SendImg(paths[r.Intn(len(paths))], wxbot.Bot.UserName, target)
	}
	if strings.Contains(FromUserName, "@@") || strings.Contains(ToUserName, "@@") {
		// ignore group message
		return
	}
	uri := "http://www.gifmiao.com/search/" + url.QueryEscape(content) + "/3"
	s, err := spider.CreateSpiderFromUrl(uri)
	if err != nil {
		logs.Error(err)
		return
	}
	srcs, _ := s.GetAttr("div.wrap>div#main>ul#waterfall>li.item>div.img_block>a>img.gifImg", "xgif")
	gif := srcs[r.Intn(len(srcs))]
	resp, err := http.Get(gif)
	if err != nil {
		wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	ls := rrstorage.CreateLocalDiskStorage("/data/gif/")
	filename := "gif-" + GetRandomStringFromNum(5) + ".gif"
	if err := ls.Save(body, filename); err != nil {
		logs.Error(err)
		return
	}
	wxbot.SendEmotion("/data/gif/" + filename, wxbot.Bot.UserName, FromUserName)
}

func delGroupText(msg map[string]interface{}) {
	logs.Debug(msg)
	content := msg["Content"].(string)
	fmt.Println(content)
	FromUserName := msg["FromUserName"].(string)
	ToUserName := msg["ToUserName"].(string)
	logs.Debug("from", FromUserName, "to", ToUserName)
	if !strings.Contains(FromUserName, "@@") && !strings.Contains(ToUserName, "@@") {
		// ignore none group message
		return
	}
	who := ""
	targetUserName := ""
	if FromUserName == wxbot.Bot.UserName {
		// from myself
		targetUserName = ToUserName
	}else {
		// from somebody else
		ss := strings.Split(content, ":")
		who = ss[0]
		content = strings.TrimPrefix(ss[1], "<br/>")
		logs.Debug("from", who)
		targetUserName = FromUserName
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	contact := wxbot.Cm.GetContactByUserName(targetUserName)
	if contact != nil {
		logs.Debug(contact)
	}else{
		logs.Error("no this contact", targetUserName)
		return
	}

	uri := "http://www.gifmiao.com/search/" + url.QueryEscape(content) + "/3"
	s, err := spider.CreateSpiderFromUrl(uri)
	if err != nil {
		logs.Error(err)
		return
	}
	srcs, _ := s.GetAttr("div.wrap>div#main>ul#waterfall>li.item>div.img_block>a>img.gifImg", "xgif")
	if len(srcs) < 1 {
		logs.Debug("112")
		return
	}
	gif := srcs[r.Intn(len(srcs))]
	resp, err := http.Get(gif)
	if err != nil {
		logs.Error(err)
		wxbot.SendText(err.Error(), wxbot.Bot.UserName, targetUserName)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	ls := rrstorage.CreateLocalDiskStorage("/data/gif/")
	filename := "gif-" + GetRandomStringFromNum(5) + ".gif"
	if err := ls.Save(body, filename); err != nil {
		logs.Error(err)
		return
	}
	logs.Debug("113")
	logs.Debug("try send")
	wxbot.SendEmotion("/data/gif/" + filename, wxbot.Bot.UserName, targetUserName)
}

func GetRandomStringFromNum(length int) string {
        bytes := []byte("0123456789")
        result := []byte{}
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        for i := 0; i < length; i++ {
                result = append(result, bytes[r.Intn(len(bytes))])
        }
        return string(result)
}

func main() {
	// add text message handler
	wxbot.HandlerRegister.Add(1, wxbot.Handler(delText))
	wxbot.HandlerRegister.Add(1, wxbot.Handler(delGroupText))
	// login
	wxbot.AutoLogin()
	// enter message loop
	wxbot.Run()
}
