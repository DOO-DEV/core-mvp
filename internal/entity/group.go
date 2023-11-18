package entity

import (
	"time"
)

type Group struct {
	MessageId    string                 `json:"messageId,omitempty"`
	AnonymousId  string                 `json:"anonymousId,omitempty"`
	UserId       string                 `json:"userId,omitempty"`
	GroupId      string                 `json:"groupId"`
	Timestamp    time.Time              `json:"timestamp,omitempty"`
	Context      *Context               `json:"context,omitempty"`
	Traits       map[string]interface{} `json:"traits,omitempty"`
	Integrations map[string]interface{} `json:"integrations,omitempty"`
}
