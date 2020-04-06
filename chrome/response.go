package chrome

const (
	WEB_ELEMENT_IDENTIFIER = "element-6066-11e4-a52e-4f735466cecf"
)

type CommonResponse struct {
	Value interface{} `json:"value"`
}

type ResponseValueError struct {
	Error string `json:"error,omitempty"`
}

type EmptyResponse struct {
	Value ResponseValueError `json:"value"`
}

type NewSessionResponse struct {
	Value SessionResponseValue `json:"value"`
}

type SessionResponseValue struct {
	Capabilities interface{} `json:"capabilities"`
	SessionId    string      `json:"sessionId"`
	Error        string      `json:"error,omitempty"`
}

type SessionsResponse struct {
	SessionId string                 `json:"sessionId"`
	Status    int                    `json:"status"`
	Value     []SessionResponseValue `json:"value"`
}

type Status struct {
	Build struct {
		Version string `json:"version"`
	} `json:"build"`
	Message string `json:"message"`
	OS      struct {
		Arch    string `json:"arch"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"os"`
	Ready bool   `json:"ready"`
	Error string `json:"error,omitempty"`
}

type StatusResponse struct {
	Value Status `json:"value"`
}

type StringResponse struct {
	Value string `json:"value"`
}

type StringsResponse struct {
	Value []string `json:"value"`
}

type BoolResponse struct {
	Value bool `json:"value"`
}

type Timeouts struct {
	Implicit int `json:"implicit"`
	PageLoad int `json:"pageLoad"`
	Script   int `json:"script"`
}

type TimeoutsResponse struct {
	Value Timeouts `json:"value"`
}

type HandleResponse struct {
	Value struct {
		Handle string `json:"handle"`
		Type   string `json:"type"`
		Error  string `json:"error,omitempty"`
	} `json:"value"`
}

type Rect struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Error  string  `json:"error,omitempty"`
}

type RectResponse struct {
	Value Rect `json:"value"`
}

type ElementResponse struct {
	Value struct {
		WebElementId string `json:"element-6066-11e4-a52e-4f735466cecf"`
		Error        string `json:"error,omitempty"`
	}
}

type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Secure   bool   `json:"secure"`
	HttpOnly bool   `json:"httpOnly"`
	Expiry   uint   `json:"expiry"`
	Error    string `json:"error,omitempty"`
}

type CookiesResponse struct {
	Value []Cookie `json:"value"`
}

type CookieResponse struct {
	Value Cookie `json:"value"`
}
