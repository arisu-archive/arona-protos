package protos

type FriendGetFriendDetailedInfoResponse struct {
	ResponsePacket
	Protocol                       Protocol
	Nickname                       string
	Level                          int64
	ClanName                       string
	Comment                        string
	FriendCount                    int64
	FriendCode                     string
	RepresentCharacterUniqueId     int64
	RepresentCharacterCostumeId    int64
	CharacterCount                 int64
	LastNormalCampaignClearStageId *int64
	LastHardCampaignClearStageId   *int64
	ArenaRanking                   *int64
	RaidRanking                    *int64
	RaidTier                       *int32
	DetailedAccountInfoDB          DetailedAccountInfoDB
	AttachmentDB                   AccountAttachmentDB
	AssistCharacterDBs             []AssistCharacterDB
}
