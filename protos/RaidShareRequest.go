package protos

type RaidShareRequest struct {
	RequestPacket
	Protocol     Protocol
	RaidServerId int64
}
