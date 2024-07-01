// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// CopyMessages Use this method to copy messages of any kind
// If some of the specified messages can't be found or copied, they are skipped
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot
// The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message
// Album grouping is kept for copied messages
// On success, an array of MessageId of the sent messages is returned.
type CopyMessages struct {
	ChatID              any     `json:"chat_id"`
	DisableNotification bool    `json:"disable_notification,omitempty"`
	FromChatID          int64   `json:"from_chat_id"`
	MessageIDs          []int64 `json:"message_ids,omitempty"`
	MessageThreadID     int64   `json:"message_thread_id,omitempty"`
	ProtectContent      bool    `json:"protect_content,omitempty"`
	RemoveCaption       bool    `json:"remove_caption,omitempty"`
}

func (entity *CopyMessages) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *CopyMessages) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity CopyMessages) MarshalJSON() ([]byte, error) {
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 CopyMessages
	return json.Marshal((x0)(entity))
}

func (CopyMessages) MethodName() string {
	return "copyMessages"
}

func (CopyMessages) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.MessageId `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeArrayOfMessageId,
		Result: x1.Result,
	}
	return &result, nil
}
