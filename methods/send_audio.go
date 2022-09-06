// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SendAudio Use this method to send audio files, if you want Telegram clients to display them in the music player
// Your audio must be in the .MP3 or .M4A format
// On success, the sent Message is returned
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	AllowSendingWithoutReply bool                      `json:"allow_sending_without_reply,omitempty"`
	Audio                    rawTypes.InputFile        `json:"audio,omitempty"`
	Caption                  string                    `json:"caption,omitempty"`
	CaptionEntities          []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID                   any                       `json:"chat_id"`
	DisableNotification      bool                      `json:"disable_notification,omitempty"`
	Duration                 int                       `json:"duration,omitempty"`
	ParseMode                string                    `json:"parse_mode,omitempty"`
	Performer                string                    `json:"performer,omitempty"`
	ProtectContent           bool                      `json:"protect_content,omitempty"`
	ReplyMarkup              any                       `json:"reply_markup,omitempty"`
	ReplyToMessageID         int64                     `json:"reply_to_message_id,omitempty"`
	Thumb                    rawTypes.InputFile        `json:"thumb,omitempty"`
	Title                    string                    `json:"title,omitempty"`
	Progress                 rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendAudio) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendAudio) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Audio.(type) {
	case types.InputBytes:
		files["audio"] = entity.Audio
		entity.Audio = nil
	}
	switch entity.Thumb.(type) {
	case types.InputBytes:
		files["thumb"] = entity.Thumb
		entity.Thumb = types.InputURL("attach://thumb")
	}
	return files
}

func (entity SendAudio) MarshalJSON() ([]byte, error) {
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
	type x0 SendAudio
	return json.Marshal((x0)(entity))
}

func (SendAudio) MethodName() string {
	return "sendAudio"
}

func (SendAudio) ParseResult(response []byte) (*rawTypes.Result, error) {
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
