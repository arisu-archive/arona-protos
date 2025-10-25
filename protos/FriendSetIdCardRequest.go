package protos

type FriendSetIdCardRequest struct {
	RequestPacket
	Protocol                   Protocol
	Comment                    string
	RepresentCharacterUniqueId int64
	EmblemId                   int64
	SearchPermission           bool
	AutoAcceptFriendRequest    bool
	ShowAccountLevel           bool
	ShowFriendCode             bool
	ShowRaidRanking            bool
	ShowArenaRanking           bool
	ShowEliminateRaidRanking   bool
	BackgroundId               int64
}
