package chrome

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"strings"
)

type WebDriver struct {
	UrlPrefix        string
	CurrentSessionId string
}

func unmarshalResponse(resBytes []byte, ret interface{}) error {
	res := EmptyResponse{}
	err := json.Unmarshal(resBytes, &res)
	if err != nil {
		return err
	}
	if res.Value.Error != "" {
		return errors.New(res.Value.Error)
	}
	if retRes, ok := ret.(EmptyResponse); ok {
		retRes.Value.Error = res.Value.Error
		return nil
	}
	if ret == nil {
		return nil
	}

	return json.Unmarshal(resBytes, ret)
}

func Get(url string, ret interface{}) error {
	fmt.Println(url)
	bodyBytes, err := ghttpclient.Get(url, nil).ReadBodyClose()
	if err != nil {
		return err
	}
	fmt.Println(string(bodyBytes))
	return unmarshalResponse(bodyBytes, ret)
}

func Post(url string, value interface{}, ret interface{}) error {
	var jsonBytes []byte
	var err error
	if value == nil {
		jsonBytes = []byte("{}")
	} else {
		jsonBytes, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}
	fmt.Println(string(jsonBytes))

	res := EmptyResponse{}
	bodyBytes, err := ghttpclient.PostJson(url, jsonBytes, nil).ReadBodyClose()
	if err != nil {
		return err
	}
	fmt.Println("body", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return err
	}
	if res.Value.Error != "" {
		return errors.New(res.Value.Error)
	}
	if retRes, ok := ret.(EmptyResponse); ok {
		retRes.Value.Error = res.Value.Error
		return nil
	}
	if ret == nil {
		return nil
	}

	return json.Unmarshal(bodyBytes, ret)
}

func Delete(url string, ret interface{}) error {
	res := EmptyResponse{}
	bodyBytes, err := ghttpclient.Delete(url, nil).ReadBodyClose()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return err
	}
	if res.Value.Error != "" {
		return errors.New(res.Value.Error)
	}
	if retRes, ok := ret.(EmptyResponse); ok {
		retRes.Value.Error = res.Value.Error
		return nil
	}
	if ret == nil {
		return nil
	}

	return json.Unmarshal(bodyBytes, ret)
}

func NewSession(urlPrefix string, capabilities *Capabilities) (*WebDriver, error) {
	api := strings.TrimRight(urlPrefix, "/") + "/session"

	params := map[string]interface{}{
		"capabilities": capabilities,
	}

	newSessionRes := NewSessionResponse{}
	err := Post(api, params, &newSessionRes)
	if err != nil {
		return nil, err
	}

	wd := WebDriver{
		UrlPrefix:        urlPrefix,
		CurrentSessionId: newSessionRes.Value.SessionId,
	}

	return &wd, nil
}

func OldSession(urlPrefix, sessionId string) *WebDriver {
	wd := WebDriver{
		UrlPrefix:        urlPrefix,
		CurrentSessionId: sessionId,
	}
	return &wd
}

func Sessions(urlPrefix string) ([]SessionResponseValue, error) {
	api := strings.TrimRight(urlPrefix, "/") + "/sessions"

	sessionsRes := SessionsResponse{}
	err := Get(api, &sessionsRes)
	if err != nil {
		return nil, err
	}

	return sessionsRes.Value, nil
}

func GetStatus(urlPrefix string) (*Status, error) {
	api := strings.TrimRight(urlPrefix, "/") + "/status"

	statusRes := StatusResponse{}
	err := Get(api, &statusRes)
	if err != nil {
		return nil, err
	}

	return &statusRes.Value, nil
}

func DeleteSession(urlPrefix, sessionId string) error {
	api := strings.TrimRight(urlPrefix, "/") + "/session/" + sessionId
	return Delete(api, nil)
}

func (wd *WebDriver) Get(path string, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Get(url, ret)
}

func (wd *WebDriver) Post(path string, value interface{}, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Post(url, value, ret)
}

func (wd *WebDriver) Delete(path string, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Delete(url, ret)
}

func (wd *WebDriver) UrlTo(url string) error {
	return wd.Post("/url", map[string]string{"url": url}, nil)
}

func (wd *WebDriver) CurrentUrl() (string, error) {
	ret := StringResponse{}
	err := wd.Get("/url", &ret)

	return ret.Value, err
}

func (wd *WebDriver) GetTimeouts() (*Timeouts, error) {
	timeoutsRes := TimeoutsResponse{}
	err := wd.Get("/timeouts", &timeoutsRes)
	if err != nil {
		return nil, err
	}

	return &timeoutsRes.Value, nil
}

func (wd *WebDriver) SetTimeouts(timeouts Timeouts) error {
	return wd.Post("/timeouts", timeouts, nil)
}

func (wd *WebDriver) Back() error {
	return wd.Post("/back", nil, nil)
}

func (wd *WebDriver) Forward() error {
	return wd.Post("/forward", nil, nil)
}

func (wd *WebDriver) Refresh() error {
	return wd.Post("/refresh", nil, nil)
}

func (wd *WebDriver) Title() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/title", &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) Window() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/window", &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) CloseWindow() error {
	return wd.Delete("/window", nil)
}

func (wd *WebDriver) SwitchWindow(handle string) error {
	return wd.Post("/window", map[string]string{"handle": handle}, nil)
}

func (wd *WebDriver) WindowHandles() ([]string, error) {
	stringsRes := StringsResponse{}
	err := wd.Get("/window/handles", &stringsRes)

	return stringsRes.Value, err
}

func (wd *WebDriver) NewWindow() (string, error) {
	stringRes := StringResponse{}
	err := wd.Post("/window/new", nil, &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) SwitchFrameId(id *int) error {
	return wd.Post("/frame", map[string]*int{"id": id}, nil)
}

func (wd *WebDriver) SwitchFrameElement(id *WebElement) error {
	return wd.Post("/frame", map[string]*WebElement{"id": id}, nil)
}

func (wd *WebDriver) SwitchParentFrame() error {
	return wd.Post("/frame/parent", nil, nil)
}

func (wd *WebDriver) WindowRect() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Get("/window/rect", &windowRectRes)

	return &windowRectRes.Value, err
}

func (wd *WebDriver) SetWindowRect(rect *Rect) (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/rect", rect, &windowRectRes)

	return &windowRectRes.Value, err
}

func (wd *WebDriver) Maximize() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/maximize", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

func (wd *WebDriver) Minimize() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/minimize", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

func (wd *WebDriver) FullScreen() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/fullscreen", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

func (wd *WebDriver) ActiveElement() (string, error) {
	elementRes := ElementResponse{}
	err := wd.Get("/element/active", &elementRes)

	return elementRes.Value.WebElementId, err
}

func (wd *WebDriver) FindElement(using, value string) (*WebElement, error) {
	webEleRes := ElementResponse{}
	err := wd.Post("/element", map[string]string{
		"using": using,
		"value": value,
	}, &webEleRes)

	fmt.Println("element res : ", webEleRes)

	webEle := WebElement{
		wd: wd,
		ID: webEleRes.Value.WebElementId,
	}

	return &webEle, err
}

func (wd *WebDriver) FindElements(using, value string) ([]WebElement, error) {
	webElesRes := struct {
		Value []WebElement `json:"value"`
	}{}
	err := wd.Post("/elements", map[string]string{
		"using": using,
		"value": value,
	}, &webElesRes)

	webEles := webElesRes.Value
	for i := range webEles {
		webEles[i].wd = wd
	}
	return webEles, err
}

func (wd *WebDriver) ScreenshotElement(elementId string) (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/element/"+elementId+"/screenshot", &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) Source() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/source", &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) ExecuteScriptSync(script string, args []interface{}, ret interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	return wd.Post("/execute/sync", map[string]interface{}{
		"script": script,
		"args":   args,
	}, ret)
}

func (wd *WebDriver) ExecuteScriptAsync(script string, args []interface{}, ret interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	return wd.Post("/execute/async", map[string]interface{}{
		"script": script,
		"args":   args,
	}, ret)
}

func (wd *WebDriver) Cookies() ([]Cookie, error) {
	cookiesRes := CookiesResponse{}
	err := wd.Get("/cookie", &cookiesRes)

	return cookiesRes.Value, err
}

func (wd *WebDriver) Cookie(name string) (*Cookie, error) {
	cookieRes := CookieResponse{}
	err := wd.Get("/cookie/"+name, &cookieRes)

	return &cookieRes.Value, err
}

func (wd *WebDriver) AddCookie(cookie Cookie) error {
	return wd.Post("/cookie", map[string]Cookie{"cookie": cookie}, nil)
}

func (wd *WebDriver) DeleteCookie(name string) error {
	return wd.Delete("/cookie/"+name, nil)
}

func (wd *WebDriver) PerformActions(actionSequences KeyActionSequences) error {
	return wd.Post("/actions", map[string]interface{}{"actions": actionSequences}, nil)
}

func (wd *WebDriver) PerformPointerActions(actionSequences PointerActionSequences) error {
	return wd.Post("/actions", map[string]interface{}{"actions": actionSequences}, nil)
}

func (wd *WebDriver) ReleaseActions() error {
	return wd.Delete("/actions", nil)
}

func (wd *WebDriver) AlertText() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/alert/text", &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) SetAlertText(text string) (string, error) {
	stringRes := StringResponse{}
	err := wd.Post("/alert/text", map[string]string{"text": text}, &stringRes)

	return stringRes.Value, err
}

func (wd *WebDriver) AlertDismiss(text string) error {
	return wd.Post("/alert/dismiss", nil, nil)
}

func (wd *WebDriver) AlertAccept(text string) error {
	return wd.Post("/alert/accept", nil, nil)
}

func (wd *WebDriver) Screenshot() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/screenshot", &stringRes)
	return stringRes.Value, err
}
