package dto

import raas "github.com/raas-app/stocks"

type MetaActionViewType string

type ResponseError struct {
	Meta *Meta `json:"meta"`
}

const (
	MetaActionViewTypeDialog MetaActionViewType = "dialog"
	MetaActionViewTypeToast  MetaActionViewType = "toast"
	MetaActionViewTypeBDUI   MetaActionViewType = "bdui"
)

type MetaAction struct {
	ViewType MetaActionViewType `json:"view_type"`
	Button   *MetaButton        `json:"button"`
	SlugID   *string            `json:"slug_id,omitempty"`
}

type MetaButton struct {
	Text string `json:"text"`
	URL  string `json:"url,omitempty"`
}

var (
	MetaActionBadRequest = &MetaAction{
		ViewType: MetaActionViewTypeToast,
		Button: &MetaButton{
			Text: "Ok",
		},
	}
)

type Meta struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	DebugID string      `json:"debug_id,omitempty"`
	Action  *MetaAction `json:"action,omitempty"`
	Reason  string      `json:"reason,omitempty"`
}

func NewMeta(code int, message, debugID string, action *MetaAction) *Meta {
	return &Meta{
		Code:    code,
		Message: message,
		DebugID: debugID,
		Action:  action,
	}
}

func NewMetaAction(action *raas.Action) *MetaAction {
	metaAction := &MetaAction{
		ViewType: MetaActionViewType(action.ViewType),
		Button: &MetaButton{
			Text: action.Button.Text,
			URL:  action.Button.URL,
		},
	}
	if action.SlugID != "" {
		metaAction.SlugID = &action.SlugID
	}

	return metaAction
}

func (e ResponseError) Error() string {
	if e.Meta.Message == "" {
		return "UNKNOWN ERROR"
	}
	return e.Meta.Message
}
