// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// InlineQueryResultCachedVideo Represents a link to a video file stored on the Telegram servers
// By default, this video file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Caption               string                `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	Description           string                `json:"description,omitempty"`
	ID                    string                `json:"id"`
	InputMessageContent   any                   `json:"input_message_content,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	Title                 string                `json:"title"`
	VideoFileID           string                `json:"video_file_id"`
}

func (entity InlineQueryResultCachedVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string                `json:"type"`
		ID                    string                `json:"id"`
		VideoFileID           string                `json:"video_file_id"`
		Title                 string                `json:"title"`
		Description           string                `json:"description,omitempty"`
		Caption               string                `json:"caption,omitempty"`
		ParseMode             string                `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent   any                   `json:"input_message_content,omitempty"`
	}{
		Type:                  "video",
		ID:                    entity.ID,
		VideoFileID:           entity.VideoFileID,
		Title:                 entity.Title,
		Description:           entity.Description,
		Caption:               entity.Caption,
		ParseMode:             entity.ParseMode,
		CaptionEntities:       entity.CaptionEntities,
		ShowCaptionAboveMedia: entity.ShowCaptionAboveMedia,
		ReplyMarkup:           entity.ReplyMarkup,
		InputMessageContent:   entity.InputMessageContent,
	}
	if !reflect.ValueOf(entity.InputMessageContent).IsNil() {
		switch entity.InputMessageContent.(type) {
		case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
			break
		default:
			return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
