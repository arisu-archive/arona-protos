package protos

type ClanKickResponse struct {
	ResponsePacket
	KickHistoryAccountIds []int64
}
