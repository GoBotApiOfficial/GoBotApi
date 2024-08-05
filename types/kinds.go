// Code AutoGenerated; DO NOT EDIT.

package types

// BACKGROUND_TYPE
const (
	TypeBackgroundTypeFill = iota
	TypeBackgroundTypeWallpaper
	TypeBackgroundTypePattern
	TypeBackgroundTypeChatTheme
)

// BACKGROUND_FILL
const (
	TypeBackgroundFillSolid = iota
	TypeBackgroundFillGradient
	TypeBackgroundFillFreeformGradient
)

// REVENUE_WITHDRAWAL_STATE
const (
	TypeRevenueWithdrawalStatePending = iota
	TypeRevenueWithdrawalStateSucceeded
	TypeRevenueWithdrawalStateFailed
)

// MESSAGE_ORIGIN
const (
	TypeMessageOriginUser = iota
	TypeMessageOriginHiddenUser
	TypeMessageOriginChat
	TypeMessageOriginChannel
)

// CHAT_BOOST_SOURCE
const (
	TypeChatBoostSourcePremium = iota
	TypeChatBoostSourceGiftCode
	TypeChatBoostSourceGiveaway
)

// CHAT_MEMBER
const (
	TypeChatMemberOwner = iota
	TypeChatMemberAdministrator
	TypeChatMemberMember
	TypeChatMemberRestricted
	TypeChatMemberLeft
	TypeChatMemberBanned
)

// MAYBE_INACCESSIBLE_MESSAGE
const (
	TypeMessage = iota
	TypeInaccessibleMessage
)

// PAID_MEDIA
const (
	TypePaidMediaPreview = iota
	TypePaidMediaPhoto
	TypePaidMediaVideo
)

// TRANSACTION_PARTNER
const (
	TypeTransactionPartnerUser = iota
	TypeTransactionPartnerFragment
	TypeTransactionPartnerTelegramAds
	TypeTransactionPartnerOther
)

// RETURN_TYPES
const (
	TypeArrayOfBotCommand = iota
	TypeArrayOfChatMember
	TypeArrayOfGameHighScore
	TypeArrayOfMessage
	TypeArrayOfMessageId
	TypeArrayOfSticker
	TypeArrayOfUpdate
	TypeBoolean
	TypeBotDescription
	TypeBotName
	TypeBotShortDescription
	TypeBusinessConnection
	TypeChatAdministratorRights
	TypeChatFullInfo
	TypeChatInviteLink
	TypeChatMember
	TypeErrorMessage
	TypeFile
	TypeForumTopic
	TypeInteger
	TypeMenuButton
	TypeMessageId
	TypePoll
	TypeSentWebAppMessage
	TypeStarTransactions
	TypeStickerSet
	TypeString
	TypeUser
	TypeUserChatBoosts
	TypeUserProfilePhotos
	TypeWebhookInfo
)
