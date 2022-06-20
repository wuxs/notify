package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/esap/wechat"
)

var cache = make(map[string]int64)

var DefaultAppId string
var DefaultAgentId int
var DefaultSecret string
var ClientCache = make(map[string]*wechat.Server)

func NewWechat(appid string, agentid int, secret string) *wechat.Server {
	cfg := &wechat.WxConfig{
		AppId:   appid,
		AgentId: agentid,
		Secret:  secret,
		AppType: 1,
	}
	app := wechat.New(cfg)
	ClientCache[appid+"|"+strconv.Itoa(agentid)] = app
	return app
}

func GetWechat(appid string, agentid int, secret string) *wechat.Server {
	if appid == "" {
		appid = DefaultAppId
	}
	if agentid == 0 || secret == "" {
		agentid = DefaultAgentId
		secret = DefaultSecret
	}
	if client, ok := ClientCache[appid+"|"+strconv.Itoa(agentid)]; ok {
		return client
	} else {
		return NewWechat(appid, agentid, secret)
	}
}

func InitRoute() {
	http.HandleFunc("/send", Send)
}

func Send(w http.ResponseWriter, r *http.Request) {
	var _appId string
	var _agentId string
	var _secret string
	var _content string
	if r.Method == "GET" {
		_appId = r.FormValue("app_id")
		_agentId = r.FormValue("agent_id")
		_secret = r.FormValue("secret")
		_content = r.FormValue("content")
		println(r.Form)
	} else if r.Method == "POST" {
		_ = r.ParseForm()
		_appId = r.PostFormValue("app_id")
		_agentId = r.PostFormValue("agent_id")
		_secret = r.PostFormValue("secret")
		_content = r.PostFormValue("content")
		println(r.PostForm)
	}
	agent_id, _ := strconv.Atoi(_agentId)
	app := GetWechat(_appId, agent_id, _secret)
	if value, ok := cache[_content]; ok {
		if value < time.Now().Unix() {
			cache[_content] = time.Now().Unix() + 30*60
			app.SendText("@all", _content)
		}
	} else {
		cache[_content] = time.Now().Unix() + 30*60
		app.SendText("@all", _content)
	}
	_, _ = w.Write([]byte("ok"))
}

func main() {
	addr := flag.String("l", ":20000", "监听地址")
	flag.StringVar(&DefaultAppId, "c", "", "企业ID") //CorpID
	flag.IntVar(&DefaultAgentId, "a", 0, "应用ID")
	flag.StringVar(&DefaultSecret, "s", "", "应用Secret")
	flag.Parse()
	InitRoute()
	fmt.Println("server running,", *addr)
	http.ListenAndServe(*addr, nil)
}
