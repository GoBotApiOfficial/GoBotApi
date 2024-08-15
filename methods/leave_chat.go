// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// LeaveChat Use this method for your bot to leave a group, supergroup or channel
// Returns True on success.
type LeaveChat struct {
	ChatID any `json:"chat_id"`
}

func (entity *LeaveChat) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *LeaveChat) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity LeaveChat) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(entity.ChatID).IsNil() {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 LeaveChat
	return json.Marshal((x0)(entity))
}

func (LeaveChat) MethodName() string {
	return "leaveChat"
}

func (LeaveChat) ParseResult(response []byte) (*rawTypes.Result, error) {
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
