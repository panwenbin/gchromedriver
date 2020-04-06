package chrome

const (
	// Mouse Actions
	CLICK        = "mouseClick"
	DOUBLE_CLICK = "mouseDoubleClick"
	MOUSE_DOWN   = "mouseButtonDown"
	MOUSE_UP     = "mouseButtonUp"
	MOVE_TO      = "mouseMoveTo"
)

const (
	// Touch Actions
	SINGLE_TAP   = "touchSingleTap"
	TOUCH_DOWN   = "touchDown"
	TOUCH_UP     = "touchUp"
	TOUCH_MOVE   = "touchMove"
	TOUCH_SCROLL = "touchScroll"
	DOUBLE_TAP   = "touchDoubleTap"
	LONG_PRESS   = "touchLongPress"
	FLICK        = "touchFlick"
)

type KeyAction struct {
	Type string `json:"type"`
	Key  string `json:"value"`
}

type PointerAction struct {
	Type     string  `json:"type"`
	Duration int     `json:"duration"`
	Button   int     `json:"button"`
	Origin   string  `json:"origin"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type KeyActionSequence struct {
	Type    string      `json:"type"`
	ID      string      `json:"id"`
	Actions []KeyAction `json:"actions"`
}

type PointerActionSequence struct {
	Type       string `json:"type"`
	Parameters struct {
		PointerType string `json:"pointerType"`
	} `json:"parameters"`
	ID      string          `json:"id"`
	Actions []PointerAction `json:"actions"`
}

type KeyActionSequences []KeyActionSequence

type PointerActionSequences []PointerActionSequence

func NewKeyActionSequences() KeyActionSequences {
	as := make(KeyActionSequences, 0)
	return as
}

func NewPointerActionSequences() PointerActionSequences {
	as := make(PointerActionSequences, 0)
	return as
}

func (as *KeyActionSequences) keyActions(actions []KeyAction) {
	justAppendActions := false
	for i := range *as {
		if (*as)[i].ID == "default keyboard" {
			justAppendActions = true
			(*as)[i].Actions = append((*as)[i].Actions, actions...)
		}
	}

	if justAppendActions == false {
		actionSequence := KeyActionSequence{
			ID:   "default keyboard",
			Type: "key",
		}
		actionSequence.Actions = append(actionSequence.Actions, actions...)
		*as = append(*as, actionSequence)
	}
}

func (as *PointerActionSequences) pointerActions(actions []PointerAction) {
	justAppendActions := false
	for i := range *as {
		if (*as)[i].ID == "mouse" {
			justAppendActions = true
			(*as)[i].Actions = append((*as)[i].Actions, actions...)
		}
	}

	if justAppendActions == false {
		actionSequence := PointerActionSequence{
			ID:   "mouse",
			Type: "pointer",
		}
		actionSequence.Parameters.PointerType = "mouse"
		actionSequence.Actions = append(actionSequence.Actions, actions...)
		*as = append(*as, actionSequence)
	}
}

func (as *KeyActionSequences) KeyDown(keys string) {
	actions := make([]KeyAction, 0)
	for _, key := range keys {
		actions = append(actions, KeyAction{
			Type: "keyDown",
			Key:  string(key),
		})
	}

	as.keyActions(actions)
}

func (as *KeyActionSequences) KeyUp(keys string) {
	actions := make([]KeyAction, 0)
	for _, key := range keys {
		actions = append(actions, KeyAction{
			Type: "keyUp",
			Key:  string(key),
		})
	}

	as.keyActions(actions)
}

func (as *PointerActionSequences) PointerDown(button int) {
	actions := make([]PointerAction, 0)
	actions = append(actions, PointerAction{
		Type:     "pointerDown",
		Duration: 0,
		Button:   button,
		Origin:   "pointer",
	})

	as.pointerActions(actions)
}

func (as *PointerActionSequences) PointerUp(button int) {
	actions := make([]PointerAction, 0)
	actions = append(actions, PointerAction{
		Type:     "pointerUp",
		Duration: 0,
		Button:   button,
		Origin:   "pointer",
	})

	as.pointerActions(actions)
}

func (as *PointerActionSequences) MouseClick() {
	as.PointerDown(LEFT_BUTTON)
	as.PointerUp(LEFT_BUTTON)
}

func (as *PointerActionSequences) MouseDoubleClick() {
	as.PointerDown(LEFT_BUTTON)
	as.PointerUp(LEFT_BUTTON)
	as.PointerDown(LEFT_BUTTON)
	as.PointerUp(LEFT_BUTTON)
}

func (as *PointerActionSequences) MouseRightClick() {
	as.PointerDown(RIGHT_BUTTON)
	as.PointerUp(RIGHT_BUTTON)
}

func (as *PointerActionSequences) MoveBy(x, y int) {
	actions := make([]PointerAction, 0)
	actions = append(actions, PointerAction{
		Type:     "pointerMove",
		Duration: 0,
		Button:   0,
		Origin:   "pointer",
		X:        float64(x),
		Y:        float64(y),
	})

	as.pointerActions(actions)
}
