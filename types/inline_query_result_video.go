// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// InlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file
// By default, this video file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Caption               string                `json:"caption,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	Description           string                `json:"description,omitempty"`
	ID                    string                `json:"id"`
	InputMessageContent   any                   `json:"input_message_content,omitempty"`
	MimeType              string                `json:"mime_type"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	VideoDuration         int64                 `json:"video_duration,omitempty"`
	VideoHeight           int64                 `json:"video_height,omitempty"`
	VideoURL              string                `json:"video_url"`
	VideoWidth            int64                 `json:"video_width,omitempty"`
}

func (entity InlineQueryResultVideo) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                  string                `json:"type"`
		ID                    string                `json:"id"`
		VideoURL              string                `json:"video_url"`
		MimeType              string                `json:"mime_type"`
		ThumbnailURL          string                `json:"thumbnail_url"`
		Title                 string                `json:"title"`
		Caption               string                `json:"caption,omitempty"`
		ParseMode             string                `json:"parse_mode,omitempty"`
		CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
		ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
		VideoWidth            int64                 `json:"video_width,omitempty"`
		VideoHeight           int64                 `json:"video_height,omitempty"`
		VideoDuration         int64                 `json:"video_duration,omitempty"`
		Description           string                `json:"description,omitempty"`
		ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent   any                   `json:"input_message_content,omitempty"`
	}{
		Type:                  "video",
		ID:                    entity.ID,
		VideoURL:              entity.VideoURL,
		MimeType:              entity.MimeType,
		ThumbnailURL:          entity.ThumbnailURL,
		Title:                 entity.Title,
		Caption:               entity.Caption,
		ParseMode:             entity.ParseMode,
		CaptionEntities:       entity.CaptionEntities,
		ShowCaptionAboveMedia: entity.ShowCaptionAboveMedia,
		VideoWidth:            entity.VideoWidth,
		VideoHeight:           entity.VideoHeight,
		VideoDuration:         entity.VideoDuration,
		Description:           entity.Description,
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
