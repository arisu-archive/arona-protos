package protos

type RaidDetailRequest struct {
	RequestPacket
	Protocol     Protocol
	RaidServerId int64
	RaidUniqueId int64
}
