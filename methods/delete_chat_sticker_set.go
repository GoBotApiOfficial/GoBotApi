// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method
// Returns True on success.
type DeleteChatStickerSet struct {
	ChatID any `json:"chat_id"`
}

func (entity *DeleteChatStickerSet) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *DeleteChatStickerSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity DeleteChatStickerSet) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 DeleteChatStickerSet
	return json.Marshal((x0)(entity))
}

func (DeleteChatStickerSet) MethodName() string {
	return "deleteChatStickerSet"
}

func (DeleteChatStickerSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
