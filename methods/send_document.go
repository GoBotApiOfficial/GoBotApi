// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SendDocument Use this method to send general files
// On success, the sent Message is returned
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	BusinessConnectionID        string                    `json:"business_connection_id,omitempty"`
	Caption                     string                    `json:"caption,omitempty"`
	CaptionEntities             []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID                      any                       `json:"chat_id"`
	DisableContentTypeDetection bool                      `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                      `json:"disable_notification,omitempty"`
	Document                    rawTypes.InputFile        `json:"document,omitempty"`
	MessageEffectID             string                    `json:"message_effect_id,omitempty"`
	MessageThreadID             int64                     `json:"message_thread_id,omitempty"`
	ParseMode                   string                    `json:"parse_mode,omitempty"`
	ProtectContent              bool                      `json:"protect_content,omitempty"`
	ReplyMarkup                 any                       `json:"reply_markup,omitempty"`
	ReplyParameters             *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Thumbnail                   rawTypes.InputFile        `json:"thumbnail,omitempty"`
	Progress                    rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendDocument) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendDocument) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Document.(type) {
	case types.InputBytes:
		files["document"] = entity.Document
		entity.Document = nil
	}
	switch entity.Thumbnail.(type) {
	case types.InputBytes:
		files["thumbnail"] = entity.Thumbnail
		entity.Thumbnail = types.InputURL("attach://thumbnail")
	}
	return files
}

func (entity SendDocument) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
		case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
			break
		default:
			return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendDocument
	return json.Marshal((x0)(entity))
}

func (SendDocument) MethodName() string {
	return "sendDocument"
}

func (SendDocument) ParseResult(response []byte) (*rawTypes.Result, error) {
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
