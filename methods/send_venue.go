// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// SendVenue Use this method to send information about a venue
// On success, the sent Message is returned.
type SendVenue struct {
	Address                  string  `json:"address"`
	AllowSendingWithoutReply bool    `json:"allow_sending_without_reply,omitempty"`
	ChatID                   any     `json:"chat_id"`
	DisableNotification      bool    `json:"disable_notification,omitempty"`
	FoursquareID             string  `json:"foursquare_id,omitempty"`
	FoursquareType           string  `json:"foursquare_type,omitempty"`
	GooglePlaceID            string  `json:"google_place_id,omitempty"`
	GooglePlaceType          string  `json:"google_place_type,omitempty"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	MessageThreadID          int64   `json:"message_thread_id,omitempty"`
	ProtectContent           bool    `json:"protect_content,omitempty"`
	ReplyMarkup              any     `json:"reply_markup,omitempty"`
	ReplyToMessageID         int64   `json:"reply_to_message_id,omitempty"`
	Title                    string  `json:"title"`
}

func (entity *SendVenue) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *SendVenue) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendVenue) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
		case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
			break
		default:
			return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	if entity.ChatID != nil {
		switch entity.ChatID.(type) {
		case int, int64, string:
			break
		default:
			return nil, fmt.Errorf("chat_id: unknown type: %T", entity.ChatID)
		}
	}
	type x0 SendVenue
	return json.Marshal((x0)(entity))
}

func (SendVenue) MethodName() string {
	return "sendVenue"
}

func (SendVenue) ParseResult(response []byte) (*rawTypes.Result, error) {
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
