// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

// SendVoice Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message
// For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio or Document)
// On success, the sent Message is returned
// Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	Caption             string                    `json:"caption,omitempty"`
	CaptionEntities     []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID              any                       `json:"chat_id"`
	DisableNotification bool                      `json:"disable_notification,omitempty"`
	Duration            int                       `json:"duration,omitempty"`
	MessageThreadID     int64                     `json:"message_thread_id,omitempty"`
	ParseMode           string                    `json:"parse_mode,omitempty"`
	ProtectContent      bool                      `json:"protect_content,omitempty"`
	ReplyMarkup         any                       `json:"reply_markup,omitempty"`
	ReplyParameters     *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Voice               rawTypes.InputFile        `json:"voice,omitempty"`
	Progress            rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendVoice) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendVoice) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Voice.(type) {
	case types.InputBytes:
		files["voice"] = entity.Voice
		entity.Voice = nil
	}
	return files
}

func (entity SendVoice) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
		case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
			break
		default:
			return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SendVoice
	return json.Marshal((x0)(entity))
}

func (SendVoice) MethodName() string {
	return "sendVoice"
}

func (SendVoice) ParseResult(response []byte) (*rawTypes.Result, error) {
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
