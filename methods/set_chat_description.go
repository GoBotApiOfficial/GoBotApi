// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SetChatDescription Use this method to change the description of a group, a supergroup or a channel
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns True on success.
type SetChatDescription struct {
	ChatID      int64  `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

func (entity *SetChatDescription) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetChatDescription) MethodName() string {
	return "setChatDescription"
}

func (SetChatDescription) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
