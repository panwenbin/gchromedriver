package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"log"
	"webdriver/chrome"
)

func their() {
	caps := selenium.Capabilities{}
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:9515")
	if err != nil {
		log.Fatalln(err)
	}
	wd.Back()
	fmt.Println(wd)
}

func myNewSession() *chrome.WebDriver {
	caps := chrome.Capabilities{
		ExcludeSwitches: []string{chrome.SWITCH_ENABLE_AUTOMATION},
	}

	wd, err := chrome.NewSession("http://127.0.0.1:9515", &caps)
	if err != nil {
		log.Fatalln(err)
	}

	return wd
}

func myOldSession(sessionId string) {
	wd := chrome.OldSession("http://127.0.0.1:9515", sessionId)
	err := wd.UrlTo("https://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//their()
	//myOldSession("d49f8011f3d811ca3c4208cf81793c0c")
	//fmt.Println(chrome.Sessions("http://127.0.0.1:9515"))
	//chrome.Status("http://127.0.0.1:9515")
	//err := chrome.DeleteSession("http://127.0.0.1:9515", "d49f8011f3d811ca3c4208cf81793c0c")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//wd := myNewSession()
	wd := chrome.OldSession("http://127.0.0.1:9515", "b60d5813e74769b22700752875e23ee0")
	//wd.UrlTo("https://www.baidu.com")
	//fmt.Println(wd.CurrentUrl())
	//chrome.Status("http://127.0.0.1:9515")
	fmt.Println(wd.CurrentSessionId)
	//fmt.Println(wd.GetTimeouts())
	//timeouts := chrome.Timeouts{
	//	Implicit: 0,
	//	PageLoad: 10000,
	//	Script:   10000,
	//}
	//wd.SetTimeouts(timeouts)
	//fmt.Println(wd.Refresh())
	//fmt.Println(wd.CloseWindow())
	//fmt.Println(wd.WindowHandles())
	//fmt.Println(wd.SetWindowRect(&chrome.Rect{
	//	Height: 768,
	//	Width:  1024,
	//	X:      -8,
	//	Y:      -8,
	//	Error:  "",
	//}))
	//fmt.Println(wd.ActiveElement())
	//fmt.Println(wd.AddCookie(chrome.Cookie{
	//	Name:     "abc",
	//	Value:    "abc",
	//	Domain:   "www.baidu.com",
	//	Path:     "/",
	//	Secure:   false,
	//	HttpOnly: false,
	//	Expiry:   1586879071,
	//	Error:    "",
	//}))
	//fmt.Println(wd.DeleteCookie("abc"))
	//ret := struct {
	//	Value string `json:"value"`
	//}{}
	//fmt.Println(wd.ExecuteScriptAsync("console.log('abcd');", nil, &ret))
	//fmt.Println(ret.Value)
	//we , err := wd.FindElement("css selector", "#kw")
	//	//if err != nil {
	//	//	log.Fatalln(err)
	//	//}
	//as := chrome.NewKeyActionSequences()
	//as.KeyDown(selenium.ControlKey + "a")
	//as.KeyUp("a" + selenium.ControlKey)
	//as.KeyDown(selenium.ControlKey + "c")
	//as.KeyUp("c" + selenium.ControlKey)
	//as.KeyDown(selenium.ControlKey + "v")
	//as.KeyUp("v" + selenium.ControlKey)
	//as.KeyDown(selenium.ControlKey + "v")
	//as.KeyUp("v" + selenium.ControlKey)
	//fmt.Println(wd.PerformKeyActions(as))
	//fmt.Println(wd.ReleaseActions())
	as := chrome.NewPointerActionSequences()
	as.MouseDoubleClick()
	fmt.Println(wd.PerformPointerActions(as))
}
