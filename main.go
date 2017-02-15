package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/songtianyi/laosj/spider"
	"github.com/songtianyi/rrframework/connector/redis"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/rrframework/storage"
	"github.com/songtianyi/wechat-go"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	if len(srcs) < 1 {
		logs.Error("no result for", content)
		return
	}
	gif := srcs[r.Intn(len(srcs))]
	resp, err := http.Get(gif)
	if err != nil {
		wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
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

	target := FromUserName
	if wxbot.Bot.UserName == FromUserName {
		target = ToUserName
	}
	wxbot.SendEmotion("/data/gif/"+filename, wxbot.Bot.UserName, target)
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
	} else {
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
	} else {
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
		logs.Error("no result for", content)
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
	wxbot.SendEmotion("/data/gif/"+filename, wxbot.Bot.UserName, targetUserName)
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

func Jiajia() {
	_, rc := rrredis.GetRedisClient("10.19.147.75:6379")
	for {
		select {
		case <-time.After(36 * time.Second):
			uri := "http://bbs.ncar.cc/thread-28825-1-1.html"
			s, err := spider.CreateSpiderFromUrl(uri)
			if err != nil {
				logs.Error(err)
				continue
			}
			srcs, _ := s.GetText("div.wp>div.wp.cl>div.pl.bm>table>tbody>tr>td.plc.ptm.pbn.vwthd>h1.ts>span")
			if len(srcs) < 1 {
				continue
			}
			logs.Debug(srcs)

			title, err := GbkToUtf8([]byte(srcs[0]))
			if err != nil {
				logs.Error(err)
				continue
			}
			h := md5.New()
			h.Write(title)
			sum := h.Sum(nil)
			sig := hex.EncodeToString(sum)

			exist, err := rc.HMExists("MEIJU:UPDATE:CACHE", sig)
			if err != nil {
				logs.Error(err)
				continue
			}
			if exist {
				continue
			}
			if err := rc.HMSet("MEIJU:UPDATE:CACHE", map[string]string{sig: "1"}); err != nil {
				logs.Error(err)
			}
			wxbot.SendText(string(title) + "\n" + uri , wxbot.Bot.UserName, wxbot.Bot.UserName)
			//jiajia := wxbot.Cm.GetContactByPinyin("jiajiashengjia")
			//if jiajia != nil {
			//      wxbot.SendText(string(title) + "<br/>" + uri , wxbot.Bot.UserName, jiajia.UserName)
			//}
		}
	}
}

func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

func main() {
	// add text message handler
	wxbot.HandlerRegister.Add(1, wxbot.Handler(delText))
	wxbot.HandlerRegister.Add(1, wxbot.Handler(delGroupText))
	// login
	wxbot.AutoLogin()
	go Jiajia()
	// enter message loop
	wxbot.Run()
}
