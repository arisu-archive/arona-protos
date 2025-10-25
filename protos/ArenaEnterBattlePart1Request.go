package protos

type ArenaEnterBattlePart1Request struct {
	RequestPacket
	Protocol                Protocol
	OpponentAccountServerId int64
	OpponentRank            int64
	OpponentIndex           int32
}
