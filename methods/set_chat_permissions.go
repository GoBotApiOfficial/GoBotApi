// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SetChatPermissions Use this method to set default chat permissions for all members
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights
// Returns True on success.
type SetChatPermissions struct {
	ChatID      int64                 `json:"chat_id"`
	Permissions types.ChatPermissions `json:"permissions"`
}

func (entity *SetChatPermissions) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetChatPermissions) MethodName() string {
	return "setChatPermissions"
}

func (SetChatPermissions) ParseResult(response []byte) (*rawTypes.Result, error) {
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
