package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/songtianyi/laosj/spider"
	"github.com/songtianyi/rrframework/config"
	"github.com/songtianyi/rrframework/connector/redis"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/rrframework/storage"
	"github.com/songtianyi/wechat-go"
	"github.com/songtianyi/wechat-go/wxweb"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Ephemeral struct {
	TaskId  int
	Source  string
	Result  string
	ImageId string
}

func dealText(msg map[string]interface{}) {
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
	if strings.Contains(strings.ToLower(content), "fache") {
		target := FromUserName
		if wxbot.Bot.UserName == FromUserName {
			target = ToUserName
		}
		err, rc := rrredis.GetRedisClient("10.19.147.75:6379")
		if err != nil {
			wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
			return
		}
		key := "TASK:PCITAGGING:" + strconv.Itoa(14)
		for count := 1; count <= 3; count++ {
			jsonStr, err := rc.LPop(key)
			if err != nil {
				wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
				return
			}
			ep := new(Ephemeral)
			if err := json.Unmarshal([]byte(jsonStr), ep); err != nil {
				wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
				return
			}

			// push back
			if _, err := rc.RPush(key, []byte(jsonStr)); err != nil {
				wxbot.SendText(err.Error(), wxbot.Bot.UserName, FromUserName)
				return
			}
			wxbot.SendImg("/data/sexx/haixiuzu/"+ep.ImageId, wxbot.Bot.UserName, target)
		}
		return
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
		wxbot.SendEmotion("/data/gif/"+filename, wxbot.Bot.UserName, target)
	}
}

func dealGroupText(msg map[string]interface{}) {
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
	if strings.Contains(strings.ToLower(content), "fache") {
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
			uris := []string{
				"http://bbs.ncar.cc/thread-28789-1-1.html",
				"http://bbs.ncar.cc/thread-29069-1-1.html",
				"http://bbs.ncar.cc/thread-28724-1-1.html",
			}
			for _, uri := range uris {
				s, err := spider.CreateSpiderFromUrl(uri)
				if err != nil {
					logs.Error(err)
					continue
				}
				srcs, _ := s.GetText("div.wp>div.wp.cl>div.pl.bm>table>tbody>tr>td.plc.ptm.pbn.vwthd>h1.ts>span")
				if len(srcs) < 1 {
					continue
				}

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
				wxbot.SendText(string(title)+"\n"+uri, wxbot.Bot.UserName, wxbot.Bot.UserName)
				//jiajia := wxbot.Cm.GetContactByQuanPin("jiajiashengjia")
				//if jiajia != nil {
				//      wxbot.SendText(string(title) + "<br/>" + uri , wxbot.Bot.UserName, jiajia.UserName)
				//}
			}
		}
	}
}

func dealImg(msg map[string]interface{}) {
	logs.Debug(msg)
	msgId := msg["MsgId"].(string)
	content := msg["Content"].(string)
	FromUserName := msg["FromUserName"].(string)
	ToUserName := msg["ToUserName"].(string)

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
	}
	contact := wxbot.Cm.GetContactByUserName(targetUserName)
	if contact != nil {
		logs.Debug(contact)
	} else {
		logs.Error("no this contact", targetUserName)
		return
	}
	b, err := wxbot.GetImg(msgId)
	if err != nil {
		logs.Error(err)
		return
	}
	res, err := Detect(msgId+".jpg", b)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(string(res))
	jc, _ := rrconfig.LoadJsonConfigFromBytes(res)
	ages, err := jc.GetSliceInt("faces.attributes.age.value")
	if err != nil {
		logs.Error(err)
		return
	}
	genders, _ := jc.GetSliceString("faces.attributes.gender.value")
	str := ""
	for i, v := range ages {
		str += genders[i] + "," + strconv.Itoa(v) + "\n"
	}
	wxbot.SendText(str, wxbot.Bot.UserName, targetUserName)
}

func getMaleUser() {
	gcs := wxbot.Cm.GetGroupContact()
	//ls := rrstorage.CreateLocalDiskStorage("/data/head/")
	for _, gc := range gcs {
		mm, err := wxbot.CreateMemberManagerFromGroupContact(gc)
		if err != nil {
			logs.Debug(err)
			continue
		}
		_ = mm.Update()
		//uris := mm.GetHeadImgUrlByGender(2)
		//for _, v := range uris {
		//	b, err := wxweb.WebWxGetIconByHeadImgUrl(wxbot.WxWebDefaultCommon, wxbot.WxWebXcg, wxbot.Cookies, v)
		//	if err != nil {
		//		logs.Debug(err)
		//		continue
		//	}
		//	filename := "head-" + GetRandomStringFromNum(5) + ".jpg"
		//	if err := ls.Save(b, filename); err != nil {
		//		logs.Error(err)
		//		continue
		//	}
		//	wxbot.SendImg("/data/head/" + filename, wxbot.Bot.UserName, wxbot.Bot.UserName)
		//}
		conts := mm.GetContactsByGender(2)
		for _, v := range conts {
			vu := make([]*wxweb.VerifyUser, 0)
			vu = append(vu, &wxweb.VerifyUser{
				Value:            v.UserName,
				VerifyUserTicket: "",
			})
			b2, err := wxweb.WebWxVerifyUser(wxbot.WxWebDefaultCommon, wxbot.WxWebXcg, wxbot.Cookies, mm.Group.NickName+" "+wxbot.Bot.NickName, vu)
			if err != nil {
				logs.Error(err)
				continue
			}
			logs.Debug(string(b2))
			time.Sleep(3 * time.Second)
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
	wxbot.HandlerRegister.Add(1, wxbot.Handler(dealText))
	wxbot.HandlerRegister.Add(1, wxbot.Handler(dealGroupText))
	wxbot.HandlerRegister.Add(3, wxbot.Handler(dealImg))
	// login
	wxbot.AutoLogin()
	go Jiajia()
	//go getMaleUser()
	// enter message loop
	wxbot.Run()
}
