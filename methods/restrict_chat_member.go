// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// RestrictChatMember Use this method to restrict a user in a supergroup
// The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights
// Pass True for all permissions to lift restrictions from a user
// Returns True on success.
type RestrictChatMember struct {
	ChatID      any                   `json:"chat_id"`
	Permissions types.ChatPermissions `json:"permissions"`
	UntilDate   int64                 `json:"until_date,omitempty"`
	UserID      int64                 `json:"user_id"`
}

func (entity *RestrictChatMember) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *RestrictChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity RestrictChatMember) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 RestrictChatMember
	return json.Marshal((x0)(entity))
}

func (RestrictChatMember) MethodName() string {
	return "restrictChatMember"
}

func (RestrictChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
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
