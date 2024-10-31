// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendVideoNote As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long
// Use this method to send video messages
// On success, the sent Message is returned.
type SendVideoNote struct {
	AllowPaidBroadcast   bool                      `json:"allow_paid_broadcast,omitempty"`
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	ChatID               any                       `json:"chat_id"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	Duration             int                       `json:"duration,omitempty"`
	Length               int                       `json:"length,omitempty"`
	MessageEffectID      string                    `json:"message_effect_id,omitempty"`
	MessageThreadID      int64                     `json:"message_thread_id,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	ReplyMarkup          any                       `json:"reply_markup,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	Thumbnail            rawTypes.InputFile        `json:"thumbnail,omitempty"`
	VideoNote            rawTypes.InputFile        `json:"video_note,omitempty"`
	Progress             rawTypes.ProgressCallable `json:"-"`
}

func (entity *SendVideoNote) ProgressCallable() rawTypes.ProgressCallable {
	return entity.Progress
}

func (entity *SendVideoNote) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Thumbnail.(type) {
	case types.InputBytes:
		files["thumbnail"] = entity.Thumbnail
		entity.Thumbnail = types.InputURL("attach://thumbnail")
	}
	switch entity.VideoNote.(type) {
	case types.InputBytes:
		files["video_note"] = entity.VideoNote
		entity.VideoNote = nil
	}
	return files
}

func (entity SendVideoNote) MarshalJSON() ([]byte, error) {
	nilCheck := func(val any) bool {
		if val == nil {
			return true
		}
		v := reflect.ValueOf(val)
		k := v.Kind()
		switch k {
		case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
			return v.IsNil()
		default:
			return false
		}
	}
	_ = nilCheck
	if nilCheck(entity.Thumbnail) {
		entity.Thumbnail = nil
	}
	if nilCheck(entity.ReplyParameters) {
		entity.ReplyParameters = nil
	}
	if nilCheck(entity.ReplyMarkup) {
		entity.ReplyMarkup = nil
	}
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
	type x0 SendVideoNote
	return json.Marshal((x0)(entity))
}

func (SendVideoNote) MethodName() string {
	return "sendVideoNote"
}

func (SendVideoNote) ParseResult(response []byte) (*rawTypes.Result, error) {
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
