// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"reflect"
)

// CreateChatSubscriptionInviteLink Use this method to create a subscription invite link for a channel chat
// The bot must have the can_invite_users administrator rights
// The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink
// Returns the new invite link as a ChatInviteLink object.
type CreateChatSubscriptionInviteLink struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	SubscriptionPeriod int    `json:"subscription_period"`
	SubscriptionPrice  int    `json:"subscription_price"`
}

func (entity *CreateChatSubscriptionInviteLink) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *CreateChatSubscriptionInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity CreateChatSubscriptionInviteLink) MarshalJSON() ([]byte, error) {
	if !reflect.ValueOf(entity.ChatID).IsNil() {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 CreateChatSubscriptionInviteLink
	return json.Marshal((x0)(entity))
}

func (CreateChatSubscriptionInviteLink) MethodName() string {
	return "createChatSubscriptionInviteLink"
}

func (CreateChatSubscriptionInviteLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatInviteLink `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}
