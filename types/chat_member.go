// Code AutoGenerated; DO NOT EDIT.

package types

// ChatMember This object contains information about one member of a chat
// Currently, the following 6 types of chat members are supported:
//  - ChatMemberOwner
//  - ChatMemberAdministrator
//  - ChatMemberMember
//  - ChatMemberRestricted
//  - ChatMemberLeft
//  - ChatMemberBanned
type ChatMember struct {
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanBeEdited           bool   `json:"can_be_edited"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanDeleteMessages     bool   `json:"can_delete_messages"`
	CanEditMessages       bool   `json:"can_edit_messages"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanManageChat         bool   `json:"can_manage_chat"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	CanManageVideoChats   bool   `json:"can_manage_video_chats"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanPostMessages       bool   `json:"can_post_messages"`
	CanPromoteMembers     bool   `json:"can_promote_members"`
	CanRestrictMembers    bool   `json:"can_restrict_members"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CustomTitle           string `json:"custom_title"`
	IsAnonymous           bool   `json:"is_anonymous"`
	IsMember              bool   `json:"is_member"`
	Status                string `json:"status"`
	UntilDate             int64  `json:"until_date"`
	User                  User   `json:"user"`
}

func (x ChatMember) Kind() int {
	switch x.Status {
	case "creator":
		return TypeChatMemberOwner
	case "administrator":
		return TypeChatMemberAdministrator
	case "member":
		return TypeChatMemberMember
	case "restricted":
		return TypeChatMemberRestricted
	case "left":
		return TypeChatMemberLeft
	case "kicked":
		return TypeChatMemberBanned
	default:
		return -1
	}
}
