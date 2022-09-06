// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SetChatPhoto Use this method to set a new profile photo for the chat
// Photos can't be changed for private chats
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns True on success.
type SetChatPhoto struct {
	ChatID   any                       `json:"chat_id"`
	Photo    rawTypes.InputFile        `json:"photo,omitempty"`
	Progress rawTypes.ProgressCallable `json:"-"`
}

func (entity *SetChatPhoto) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SetChatPhoto) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Photo.(type) {
	case types.InputBytes:
		files["photo"] = entity.Photo
		entity.Photo = nil
	}
	return files
}

func (entity SetChatPhoto) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SetChatPhoto
	return json.Marshal((x0)(entity))
}

func (SetChatPhoto) MethodName() string {
	return "setChatPhoto"
}

func (SetChatPhoto) ParseResult(response []byte) (*rawTypes.Result, error) {
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
