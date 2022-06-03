// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
)

// InlineQueryResultCachedGif Represents a link to an animated GIF file stored on the Telegram servers
// By default, this animated GIF file will be sent by the user with an optional caption
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Caption             string                `json:"caption,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	GifFileID           string                `json:"gif_file_id"`
	ID                  string                `json:"id"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Title               string                `json:"title,omitempty"`
}

func (entity InlineQueryResultCachedGif) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type                string                `json:"type"`
		ID                  string                `json:"id"`
		GifFileID           string                `json:"gif_file_id"`
		Title               string                `json:"title,omitempty"`
		Caption             string                `json:"caption,omitempty"`
		ParseMode           string                `json:"parse_mode,omitempty"`
		CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
		ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	}{
		Type:                "gif",
		ID:                  entity.ID,
		GifFileID:           entity.GifFileID,
		Title:               entity.Title,
		Caption:             entity.Caption,
		ParseMode:           entity.ParseMode,
		CaptionEntities:     entity.CaptionEntities,
		ReplyMarkup:         entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
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
