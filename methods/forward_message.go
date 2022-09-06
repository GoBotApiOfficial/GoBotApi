// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// ForwardMessage Use this method to forward messages of any kind
// Service messages can't be forwarded
// On success, the sent Message is returned.
type ForwardMessage struct {
	ChatID              any   `json:"chat_id"`
	DisableNotification bool  `json:"disable_notification,omitempty"`
	FromChatID          int64 `json:"from_chat_id"`
	MessageID           int64 `json:"message_id"`
	ProtectContent      bool  `json:"protect_content,omitempty"`
}

func (entity *ForwardMessage) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *ForwardMessage) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity ForwardMessage) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 ForwardMessage
	return json.Marshal((x0)(entity))
}

func (ForwardMessage) MethodName() string {
	return "forwardMessage"
}

func (ForwardMessage) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.Message `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeMessage,
		Result: x1.Result,
	}
	return &result, nil
}
