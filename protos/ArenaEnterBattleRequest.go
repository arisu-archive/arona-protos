package protos

type ArenaEnterBattleRequest struct {
	RequestPacket
	Protocol                Protocol
	OpponentAccountServerId int64
	OpponentIndex           int64
}
