package chrome

type WebElement struct {
	wd    *WebDriver
	ID    string `json:"id"`
	Error string `json:"error,omitempty"`
}

func (we *WebElement) FindElement(using, value string) (*WebElement, error) {
	webEleRes := struct {
		Value WebElement `json:"value"`
	}{}
	err := we.wd.Post("/element/"+we.ID+"/element", map[string]string{
		"using": using,
		"value": value,
	}, &webEleRes)

	webEle := webEleRes.Value
	webEle.wd = we.wd

	return &webEle, err
}

func (we *WebElement) FindElements(using, value string) ([]WebElement, error) {
	webElesRes := struct {
		Value []WebElement `json:"value"`
	}{}
	err := we.wd.Post("/element/"+we.ID+"/elements", map[string]string{
		"using": using,
		"value": value,
	}, &webElesRes)

	webEles := webElesRes.Value
	for i := range webEles {
		webEles[i].wd = we.wd
	}
	return webEles, err
}

func (we *WebElement) IsSelected() (bool, error) {
	boolRes := BoolResponse{}
	err := we.wd.Get("/element/"+we.ID+"/selected", &boolRes)

	return boolRes.Value, err
}

func (we *WebElement) GetAttribute(name string) (string, error) {
	stringRes := StringResponse{}
	err := we.wd.Get("/element/"+we.ID+"/attribute/"+name, &stringRes)

	return stringRes.Value, err
}

func (we *WebElement) GetProperty(name string) (string, error) {
	stringRes := StringResponse{}
	err := we.wd.Get("/element/"+we.ID+"/property/"+name, &stringRes)

	return stringRes.Value, err
}

func (we *WebElement) GetCss(name string) (string, error) {
	stringRes := StringResponse{}
	err := we.wd.Get("/element/"+we.ID+"/css/"+name, &stringRes)

	return stringRes.Value, err
}

func (we *WebElement) Text() (string, error) {
	stringRes := StringResponse{}
	err := we.wd.Get("/element/"+we.ID+"/text", &stringRes)

	return stringRes.Value, err
}

func (we *WebElement) TagName() (string, error) {
	stringRes := StringResponse{}
	err := we.wd.Get("/element/"+we.ID+"/name", &stringRes)

	return stringRes.Value, err
}

func (we *WebElement) Rect() (*Rect, error) {
	rectRes := RectResponse{}
	err := we.wd.Get("/element/"+we.ID+"/rect", &rectRes)
	return &rectRes.Value, err
}

func (we *WebElement) IsEnabled() (bool, error) {
	boolRes := BoolResponse{}
	err := we.wd.Get("/element/"+we.ID+"/enabled", &boolRes)

	return boolRes.Value, err
}

// Mouse buttons.
const (
	LEFT_BUTTON = iota
	MIDDLE_BUTTON
	RIGHT_BUTTON
)

func (we *WebElement) Click(btn int) error {
	return we.wd.Post("/element/"+we.ID+"/click", map[string]int{"button": btn}, nil)
}

func (we *WebElement) Clear() error {
	return we.wd.Post("/element/"+we.ID+"/clear", nil, nil)
}

func (we *WebElement) SendKeys(keys string) error {
	return we.wd.Post("/element/"+we.ID+"/value", map[string]string{"text": keys}, nil)
}
