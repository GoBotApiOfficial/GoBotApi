// Code AutoGenerated; DO NOT EDIT.

package types

// Audio Represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	Duration     int        `json:"duration"`
	FileID       string     `json:"file_id"`
	FileName     string     `json:"file_name,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	FileUniqueID string     `json:"file_unique_id"`
	MimeType     string     `json:"mime_type,omitempty"`
	Performer    string     `json:"performer,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	Title        string     `json:"title,omitempty"`
}
