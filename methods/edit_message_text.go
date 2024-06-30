// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// EditMessageText Use this method to edit text and game messages
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned
// Note that business messages that were not sent by the bot and do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageText struct {
	BusinessConnectionID string                      `json:"business_connection_id,omitempty"`
	ChatID               any                         `json:"chat_id,omitempty"`
	Entities             []types.MessageEntity       `json:"entities,omitempty"`
	InlineMessageID      string                      `json:"inline_message_id,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions   `json:"link_preview_options,omitempty"`
	MessageID            int64                       `json:"message_id,omitempty"`
	ParseMode            string                      `json:"parse_mode,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Text                 string                      `json:"text"`
}

func (entity *EditMessageText) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *EditMessageText) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity EditMessageText) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 EditMessageText
	return json.Marshal((x0)(entity))
}

func (EditMessageText) MethodName() string {
	return "editMessageText"
}

func (EditMessageText) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x0 struct {
		Result bool `json:"result"`
	}
	_ = json.Unmarshal(response, &x0)
	if x0.Result {
		result := rawTypes.Result{
			Kind:   types.TypeBoolean,
			Result: true,
		}
		return &result, nil
	} else {
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
}
