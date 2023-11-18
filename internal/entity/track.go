package entity

import (
	"time"
)

type Track struct {
	MessageId    string                 `json:"messageId,omitempty"`
	AnonymousId  string                 `json:"anonymousId,omitempty"`
	UserId       string                 `json:"userId,omitempty"`
	Event        string                 `json:"event"`
	Timestamp    time.Time              `json:"timestamp,omitempty"`
	Context      *Context               `json:"context,omitempty"`
	Properties   map[string]interface{} `json:"properties,omitempty"`
	Integrations map[string]interface{} `json:"integrations,omitempty"`
}
