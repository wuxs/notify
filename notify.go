package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/esap/wechat"
)

var cache = make(map[string]int64)

var DefaultAppId string
var DefaultAgentId string
var DefaultSecret string
var DefaultListen string
var ClientCache = make(map[string]*wechat.Server)

func init() {
	DefaultAppId = os.Getenv("APP_ID")
	DefaultAgentId = os.Getenv("AGENT_ID")
	DefaultSecret = os.Getenv("SECRET")
	DefaultListen = os.Getenv("LISTEN")
}
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

func GetWechat(appid string, agentid string, secret string) *wechat.Server {
	if appid == "" {
		appid = DefaultAppId
	}
	if agentid == "" || secret == "" {
		agentid = DefaultAgentId
		secret = DefaultSecret
	}
	if client, ok := ClientCache[appid+"|"+agentid]; ok {
		return client
	} else {
		agentId, err := strconv.Atoi(agentid)
		if err != nil {
			panic(err)
		}
		return NewWechat(appid, agentId, secret)
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
	app := GetWechat(_appId, _agentId, _secret)
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
	InitRoute()
	fmt.Println("server running,", DefaultListen)
	http.ListenAndServe(DefaultListen, nil)
}
