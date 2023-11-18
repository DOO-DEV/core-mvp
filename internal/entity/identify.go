package entity

import (
	"time"
)

type Identify struct {
	Type string `json:"type,omitempty"`

	MessageId    string                 `json:"messageId,omitempty"`
	AnonymousId  string                 `json:"anonymousId,omitempty"`
	UserId       string                 `json:"userId,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Timestamp    time.Time              `json:"timestamp,omitempty"`
	Context      *Context               `json:"context,omitempty"`
	Integrations map[string]interface{} `json:"integrations,omitempty"`
}
