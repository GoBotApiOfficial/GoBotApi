// Code AutoGenerated; DO NOT EDIT.

package types

// ChatBoostSource This object describes the source of a chat boost
// It can be one of
//   - ChatBoostSourcePremium
//   - ChatBoostSourceGiftCode
//   - ChatBoostSourceGiveaway
type ChatBoostSource struct {
	GiveawayMessageID int64  `json:"giveaway_message_id"`
	IsUnclaimed       bool   `json:"is_unclaimed"`
	PrizeStarCount    int    `json:"prize_star_count"`
	Source            string `json:"source"`
	User              User   `json:"user"`
}

func (x ChatBoostSource) Kind() int {
	switch x.Source {
	case "premium":
		return TypeChatBoostSourcePremium
	case "gift_code":
		return TypeChatBoostSourceGiftCode
	case "giveaway":
		return TypeChatBoostSourceGiveaway
	default:
		return -1
	}
}
