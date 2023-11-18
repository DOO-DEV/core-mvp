package entity

import (
	"time"
)

type Alias struct {
	MessageId    string                 `json:"messageId,omitempty"`
	PreviousId   string                 `json:"previousId"`
	UserId       string                 `json:"userId"`
	Timestamp    time.Time              `json:"timestamp,omitempty"`
	Context      *Context               `json:"context,omitempty"`
	Integrations map[string]interface{} `json:"integrations,omitempty"`
}
