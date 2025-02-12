// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SendGift Sends a gift to the given user or channel chat
// The gift can't be converted to Telegram Stars by the receiver
// Returns True on success.
type SendGift struct {
	ChatID        any                   `json:"chat_id,omitempty"`
	GiftID        string                `json:"gift_id"`
	PayForUpgrade bool                  `json:"pay_for_upgrade,omitempty"`
	Text          string                `json:"text,omitempty"`
	TextEntities  []types.MessageEntity `json:"text_entities,omitempty"`
	TextParseMode string                `json:"text_parse_mode,omitempty"`
	UserID        int64                 `json:"user_id,omitempty"`
}

func (entity *SendGift) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendGift) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendGift) MarshalJSON() ([]byte, error) {
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
	if nilCheck(entity.ChatID) {
		entity.ChatID = nil
	}
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SendGift
	return json.Marshal((x0)(entity))
}

func (SendGift) MethodName() string {
	return "sendGift"
}

func (SendGift) ParseResult(response []byte) (*rawTypes.Result, error) {
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
