// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultVenue Represents a venue
// By default, the venue will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April, 2016
// Older clients will ignore them.
type InlineQueryResultVenue struct {
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceID       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	ID                  string                `json:"id"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int64                 `json:"thumb_width,omitempty"`
	Title               string                `json:"title"`
}

func (entity InlineQueryResultVenue) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		Latitude            float64               `json:"latitude"`
		Longitude           float64               `json:"longitude"`
		Title               string                `json:"title"`
		Address             string                `json:"address"`
		FoursquareID        string                `json:"foursquare_id,omitempty"`
		FoursquareType      string                `json:"foursquare_type,omitempty"`
		GooglePlaceID       string                `json:"google_place_id,omitempty"`
		GooglePlaceType     string                `json:"google_place_type,omitempty"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{}           `json:"input_message_content,omitempty"`
		ThumbURL            string                `json:"thumb_url,omitempty"`
		ThumbWidth          int64                 `json:"thumb_width,omitempty"`
		ThumbHeight         int                   `json:"thumb_height,omitempty"`
	}{
		Type:                "venue",
		ID:                  entity.ID,
		Latitude:            entity.Latitude,
		Longitude:           entity.Longitude,
		Title:               entity.Title,
		Address:             entity.Address,
		FoursquareID:        entity.FoursquareID,
		FoursquareType:      entity.FoursquareType,
		GooglePlaceID:       entity.GooglePlaceID,
		GooglePlaceType:     entity.GooglePlaceType,
		ReplyMarkup:         entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
		ThumbURL:            entity.ThumbURL,
		ThumbWidth:          entity.ThumbWidth,
		ThumbHeight:         entity.ThumbHeight,
	}
	if entity.InputMessageContent != nil {
		switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
