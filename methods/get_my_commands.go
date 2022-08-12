// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// GetMyCommands Use this method to get the current list of the bot's commands for the given scope and user language
// Returns an Array of BotCommand objects
// If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	LanguageCode string                 `json:"language_code,omitempty"`
	Scope        *types.BotCommandScope `json:"scope,omitempty"`
}

func (entity *GetMyCommands) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetMyCommands) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetMyCommands) MethodName() string {
	return "getMyCommands"
}

func (GetMyCommands) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.BotCommand `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeArrayOfBotCommand,
		Result: x1.Result,
	}
	return &result, nil
}
