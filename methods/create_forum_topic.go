// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// CreateForumTopic Use this method to create a topic in a forum supergroup chat
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights
// Returns information about the created topic as a ForumTopic object.
type CreateForumTopic struct {
	ChatID            any    `json:"chat_id"`
	IconColor         int    `json:"icon_color,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	Name              string `json:"name"`
}

func (entity *CreateForumTopic) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *CreateForumTopic) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity CreateForumTopic) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(entity.ChatID).IsNil() {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 CreateForumTopic
	return json.Marshal((x0)(entity))
}

func (CreateForumTopic) MethodName() string {
	return "createForumTopic"
}

func (CreateForumTopic) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ForumTopic `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeForumTopic,
		Result: x1.Result,
	}
	return &result, nil
}
