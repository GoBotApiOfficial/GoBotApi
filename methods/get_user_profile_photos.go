// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// GetUserProfilePhotos Use this method to get a list of profile pictures for a user
// Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
	Limit  int   `json:"limit,omitempty"`
	Offset int   `json:"offset,omitempty"`
	UserID int64 `json:"user_id"`
}

func (entity *GetUserProfilePhotos) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetUserProfilePhotos) MethodName() string {
	return "getUserProfilePhotos"
}

func (GetUserProfilePhotos) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.UserProfilePhotos `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeUserProfilePhotos,
		Result: x1.Result,
	}
	return &result, nil
}
