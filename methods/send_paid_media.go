// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// SendPaidMedia Use this method to send paid media to channel chats
// On success, the sent Message is returned.
type SendPaidMedia struct {
	Caption               string                    `json:"caption,omitempty"`
	CaptionEntities       []types.MessageEntity     `json:"caption_entities,omitempty"`
	ChatID                any                       `json:"chat_id"`
	DisableNotification   bool                      `json:"disable_notification,omitempty"`
	Media                 []types.InputPaidMedia    `json:"media,omitempty"`
	ParseMode             string                    `json:"parse_mode,omitempty"`
	ProtectContent        bool                      `json:"protect_content,omitempty"`
	ReplyMarkup           any                       `json:"reply_markup,omitempty"`
	ReplyParameters       *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	ShowCaptionAboveMedia bool                      `json:"show_caption_above_media,omitempty"`
	StarCount             int                       `json:"star_count"`
	Progress              rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendPaidMedia) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendPaidMedia) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	for i, x0 := range entity.Media {
		x1 := x0.(rawTypes.InputMediaFiles).Files()
		for k, v := range x1 {
			var attachName string
			if k == "thumbnail" {
				attachName = fmt.Sprintf("file-%d-thumbnail", i)
				x0.SetAttachmentThumb(attachName)
			} else {
				attachName = fmt.Sprintf("file-%d", i)
				x0.SetAttachment(attachName)
			}
			files[attachName] = v
		}
	}
	return files
}

func (entity SendPaidMedia) MarshalJSON() ([]byte, error) {
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
	type x0 SendPaidMedia
	return json.Marshal((x0)(entity))
}

func (SendPaidMedia) MethodName() string {
	return "sendPaidMedia"
}

func (SendPaidMedia) ParseResult(response []byte) (*rawTypes.Result, error) {
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
