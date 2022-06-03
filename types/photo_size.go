// Code AutoGenerated; DO NOT EDIT.

package types

// PhotoSize Represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileSize     int    `json:"file_size,omitempty"`
	FileUniqueID string `json:"file_unique_id"`
	Height       int    `json:"height"`
	Width        int64  `json:"width"`
}
