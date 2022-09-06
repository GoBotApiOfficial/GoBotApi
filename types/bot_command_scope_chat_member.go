// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// BotCommandScopeChatMember Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	ChatID any    `json:"chat_id"`
	Type   string `json:"type"`
	UserID int64  `json:"user_id"`
}

func (entity BotCommandScopeChatMember) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 BotCommandScopeChatMember
	return json.Marshal((x0)(entity))
}
