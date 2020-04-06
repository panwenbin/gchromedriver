package chrome

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/panwenbin/ghttpclient"
	"strings"
)

// WebDriver indicates a web driver session
type WebDriver struct {
	UrlPrefix        string
	CurrentSessionId string
}

// unmarshalResponse unmarshal the response body bytes to a struct
// if an error is detected in response, returns an error
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

// Get sends a get request to rest api, and unmarshal the response
func Get(url string, ret interface{}) error {
	bodyBytes, err := ghttpclient.Get(url, nil).ReadBodyClose()
	if err != nil {
		return err
	}
	return unmarshalResponse(bodyBytes, ret)
}

// Post sends a post request to rest api, and unmarshal the response
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

// Delete sends a delete request to rest api
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

// NewSession creates a new session from url prefix and capabilities
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

// OldSession reuses an old session by url prefix and session id
func OldSession(urlPrefix, sessionId string) *WebDriver {
	wd := WebDriver{
		UrlPrefix:        urlPrefix,
		CurrentSessionId: sessionId,
	}
	return &wd
}

// Sessions lists all sessions
func Sessions(urlPrefix string) ([]SessionResponseValue, error) {
	api := strings.TrimRight(urlPrefix, "/") + "/sessions"

	sessionsRes := SessionsResponse{}
	err := Get(api, &sessionsRes)
	if err != nil {
		return nil, err
	}

	return sessionsRes.Value, nil
}

// GetStatus get the status of driver
func GetStatus(urlPrefix string) (*Status, error) {
	api := strings.TrimRight(urlPrefix, "/") + "/status"

	statusRes := StatusResponse{}
	err := Get(api, &statusRes)
	if err != nil {
		return nil, err
	}

	return &statusRes.Value, nil
}

// DeleteSession deletes a session by url prefix and session id
func DeleteSession(urlPrefix, sessionId string) error {
	api := strings.TrimRight(urlPrefix, "/") + "/session/" + sessionId
	return Delete(api, nil)
}

// Get sends a get request to rest api, and unmarshal the response
func (wd *WebDriver) Get(path string, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Get(url, ret)
}

// Post sends a post request to rest api, and unmarshal the response
func (wd *WebDriver) Post(path string, value interface{}, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Post(url, value, ret)
}

// Delete sends a delete request to rest api
func (wd *WebDriver) Delete(path string, ret interface{}) error {
	url := strings.TrimRight(wd.UrlPrefix, "/") + "/session/" + wd.CurrentSessionId + path
	return Delete(url, ret)
}

// UrlTo navigates to the url
func (wd *WebDriver) UrlTo(url string) error {
	return wd.Post("/url", map[string]string{"url": url}, nil)
}

// CurrentUrl gets the current url
func (wd *WebDriver) CurrentUrl() (string, error) {
	ret := StringResponse{}
	err := wd.Get("/url", &ret)

	return ret.Value, err
}

// GetTimeouts gets timeout settings
func (wd *WebDriver) GetTimeouts() (*Timeouts, error) {
	timeoutsRes := TimeoutsResponse{}
	err := wd.Get("/timeouts", &timeoutsRes)
	if err != nil {
		return nil, err
	}

	return &timeoutsRes.Value, nil
}

// SetTimeouts sets timeout settings
func (wd *WebDriver) SetTimeouts(timeouts Timeouts) error {
	return wd.Post("/timeouts", timeouts, nil)
}

// Back navigates back
func (wd *WebDriver) Back() error {
	return wd.Post("/back", nil, nil)
}

// Forward navigates forward
func (wd *WebDriver) Forward() error {
	return wd.Post("/forward", nil, nil)
}

// Refresh refresh the current page
func (wd *WebDriver) Refresh() error {
	return wd.Post("/refresh", nil, nil)
}

// Title gets the current title of the page
func (wd *WebDriver) Title() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/title", &stringRes)

	return stringRes.Value, err
}

// Window gets the current window handle
func (wd *WebDriver) Window() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/window", &stringRes)

	return stringRes.Value, err
}

// CloseWindow closes the current window or tab
func (wd *WebDriver) CloseWindow() error {
	return wd.Delete("/window", nil)
}

// SwitchWindow switches to a window handle as new current handle
func (wd *WebDriver) SwitchWindow(handle string) error {
	return wd.Post("/window", map[string]string{"handle": handle}, nil)
}

// WindowHandles lists all window handles
func (wd *WebDriver) WindowHandles() ([]string, error) {
	stringsRes := StringsResponse{}
	err := wd.Get("/window/handles", &stringsRes)

	return stringsRes.Value, err
}

// NewWindow creates a new window or tab
func (wd *WebDriver) NewWindow() (string, error) {
	stringRes := StringResponse{}
	err := wd.Post("/window/new", nil, &stringRes)

	return stringRes.Value, err
}

// SwitchFrameId switches to a frame by frame id
func (wd *WebDriver) SwitchFrameId(id *int) error {
	return wd.Post("/frame", map[string]*int{"id": id}, nil)
}

// SwitchFrameElement switches to a frame by a web element id
func (wd *WebDriver) SwitchFrameElement(id *WebElement) error {
	return wd.Post("/frame", map[string]*WebElement{"id": id}, nil)
}

// SwitchParentFrame switches to parent frame
func (wd *WebDriver) SwitchParentFrame() error {
	return wd.Post("/frame/parent", nil, nil)
}

// WindowRect get the rect options of the current window
func (wd *WebDriver) WindowRect() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Get("/window/rect", &windowRectRes)

	return &windowRectRes.Value, err
}

// SetWindowRect sets the rect options to the current window.
// which can resize window and move to position
func (wd *WebDriver) SetWindowRect(rect *Rect) (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/rect", rect, &windowRectRes)

	return &windowRectRes.Value, err
}

// Maximize Maximizes the window
func (wd *WebDriver) Maximize() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/maximize", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

// Minimize Minimizes the window
func (wd *WebDriver) Minimize() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/minimize", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

// FullScreen set the window fullscreen
func (wd *WebDriver) FullScreen() (*Rect, error) {
	windowRectRes := RectResponse{}
	err := wd.Post("/window/fullscreen", nil, &windowRectRes)
	return &windowRectRes.Value, err
}

// ActiveElement gets the active element id
func (wd *WebDriver) ActiveElement() (string, error) {
	elementRes := ElementResponse{}
	err := wd.Get("/element/active", &elementRes)

	return elementRes.Value.WebElementId, err
}

// FindElement finds a web element by css selector, xpath, link text
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

// FindElements finds web elements by css selector, xpath, link text
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

// ScreenshotElement takes a screenshot on the element
func (wd *WebDriver) ScreenshotElement(elementId string) (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/element/"+elementId+"/screenshot", &stringRes)

	return stringRes.Value, err
}

// Source gets the page source
func (wd *WebDriver) Source() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/source", &stringRes)

	return stringRes.Value, err
}

// ExecuteScriptSync executes a script sync and gets the return.
// a return expression is needed to get the return
func (wd *WebDriver) ExecuteScriptSync(script string, args []interface{}, ret interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	return wd.Post("/execute/sync", map[string]interface{}{
		"script": script,
		"args":   args,
	}, ret)
}

// ExecuteScriptAsync executes a script async
func (wd *WebDriver) ExecuteScriptAsync(script string, args []interface{}, ret interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	return wd.Post("/execute/async", map[string]interface{}{
		"script": script,
		"args":   args,
	}, ret)
}

// Cookies gets all cookies on the page for the current url
func (wd *WebDriver) Cookies() ([]Cookie, error) {
	cookiesRes := CookiesResponse{}
	err := wd.Get("/cookie", &cookiesRes)

	return cookiesRes.Value, err
}

// Cookies gets a cookies on the page for the current url by name
func (wd *WebDriver) Cookie(name string) (*Cookie, error) {
	cookieRes := CookieResponse{}
	err := wd.Get("/cookie/"+name, &cookieRes)

	return &cookieRes.Value, err
}

// AddCookie adds a new cookie to the page
// only cookie for the current url is accepted
func (wd *WebDriver) AddCookie(cookie Cookie) error {
	return wd.Post("/cookie", map[string]Cookie{"cookie": cookie}, nil)
}

// DeleteCookie deletes a cookie on the page for the current url by name
func (wd *WebDriver) DeleteCookie(name string) error {
	return wd.Delete("/cookie/"+name, nil)
}

// PerformKeyActions performs low level key actions
func (wd *WebDriver) PerformKeyActions(actionSequences KeyActionSequences) error {
	return wd.Post("/actions", map[string]interface{}{"actions": actionSequences}, nil)
}

// PerformPointerActions performs low level pointer actions
func (wd *WebDriver) PerformPointerActions(actionSequences PointerActionSequences) error {
	return wd.Post("/actions", map[string]interface{}{"actions": actionSequences}, nil)
}

// ReleaseActions releases actions which previous performed
func (wd *WebDriver) ReleaseActions() error {
	return wd.Delete("/actions", nil)
}

// AlertText gets the text on the alert
func (wd *WebDriver) AlertText() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/alert/text", &stringRes)

	return stringRes.Value, err
}

// SetAlertText sets the text on the alert
func (wd *WebDriver) SetAlertText(text string) (string, error) {
	stringRes := StringResponse{}
	err := wd.Post("/alert/text", map[string]string{"text": text}, &stringRes)

	return stringRes.Value, err
}

// AlertDismiss dismisses the alert
func (wd *WebDriver) AlertDismiss(text string) error {
	return wd.Post("/alert/dismiss", nil, nil)
}

// AlertAccept accepts the alert
func (wd *WebDriver) AlertAccept(text string) error {
	return wd.Post("/alert/accept", nil, nil)
}

// Screenshot takes a screenshot for the page
func (wd *WebDriver) Screenshot() (string, error) {
	stringRes := StringResponse{}
	err := wd.Get("/screenshot", &stringRes)
	return stringRes.Value, err
}
