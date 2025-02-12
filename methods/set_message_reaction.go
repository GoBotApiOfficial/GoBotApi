// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// SetMessageReaction Use this method to change the chosen reactions on a message
// Service messages of some types can't be reacted to
// Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel
// Bots can't use paid reactions
// Returns True on success.
type SetMessageReaction struct {
	ChatID    any                  `json:"chat_id"`
	IsBig     bool                 `json:"is_big,omitempty"`
	MessageID int64                `json:"message_id"`
	Reaction  []types.ReactionType `json:"reaction,omitempty"`
}

func (entity *SetMessageReaction) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SetMessageReaction) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SetMessageReaction) MarshalJSON() ([]byte, error) {
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
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SetMessageReaction
	return json.Marshal((x0)(entity))
}

func (SetMessageReaction) MethodName() string {
	return "setMessageReaction"
}

func (SetMessageReaction) ParseResult(response []byte) (*rawTypes.Result, error) {
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
