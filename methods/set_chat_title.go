// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SetChatTitle Use this method to change the title of a chat
// Titles can't be changed for private chats
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns True on success.
type SetChatTitle struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

func (entity *SetChatTitle) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetChatTitle) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SetChatTitle) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(entity.ChatID).IsNil() {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SetChatTitle
	return json.Marshal((x0)(entity))
}

func (SetChatTitle) MethodName() string {
	return "setChatTitle"
}

func (SetChatTitle) ParseResult(response []byte) (*rawTypes.Result, error) {
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
