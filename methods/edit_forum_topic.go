// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// EditForumTopic Use this method to edit name and icon of a topic in a forum supergroup chat
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic
// Returns True on success.
type EditForumTopic struct {
	ChatID            any    `json:"chat_id"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	MessageThreadID   int64  `json:"message_thread_id"`
	Name              string `json:"name,omitempty"`
}

func (entity *EditForumTopic) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *EditForumTopic) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity EditForumTopic) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(entity.ChatID).IsNil() {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 EditForumTopic
	return json.Marshal((x0)(entity))
}

func (EditForumTopic) MethodName() string {
	return "editForumTopic"
}

func (EditForumTopic) ParseResult(response []byte) (*rawTypes.Result, error) {
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
