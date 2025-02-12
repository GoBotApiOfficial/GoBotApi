// Code AutoGenerated; DO NOT EDIT.

package types

// MESSAGE_ORIGIN
const (
	TypeMessageOriginUser = iota
	TypeMessageOriginHiddenUser
	TypeMessageOriginChat
	TypeMessageOriginChannel
)

// BACKGROUND_TYPE
const (
	TypeBackgroundTypeFill = iota
	TypeBackgroundTypeWallpaper
	TypeBackgroundTypePattern
	TypeBackgroundTypeChatTheme
)

// CHAT_BOOST_SOURCE
const (
	TypeChatBoostSourcePremium = iota
	TypeChatBoostSourceGiftCode
	TypeChatBoostSourceGiveaway
)

// MAYBE_INACCESSIBLE_MESSAGE
const (
	TypeMessage = iota
	TypeInaccessibleMessage
)

// BACKGROUND_FILL
const (
	TypeBackgroundFillSolid = iota
	TypeBackgroundFillGradient
	TypeBackgroundFillFreeformGradient
)

// TRANSACTION_PARTNER
const (
	TypeTransactionPartnerUser = iota
	TypeTransactionPartnerChat
	TypeTransactionPartnerAffiliateProgram
	TypeTransactionPartnerFragment
	TypeTransactionPartnerTelegramAds
	TypeTransactionPartnerTelegramApi
	TypeTransactionPartnerOther
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

// PAID_MEDIA
const (
	TypePaidMediaPreview = iota
	TypePaidMediaPhoto
	TypePaidMediaVideo
)

// REVENUE_WITHDRAWAL_STATE
const (
	TypeRevenueWithdrawalStatePending = iota
	TypeRevenueWithdrawalStateSucceeded
	TypeRevenueWithdrawalStateFailed
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
	TypeGifts
	TypeInteger
	TypeMenuButton
	TypeMessageId
	TypePoll
	TypePreparedInlineMessage
	TypeSentWebAppMessage
	TypeStarTransactions
	TypeStickerSet
	TypeString
	TypeUser
	TypeUserChatBoosts
	TypeUserProfilePhotos
	TypeWebhookInfo
)
