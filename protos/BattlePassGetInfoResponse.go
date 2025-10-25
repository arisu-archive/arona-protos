package protos

type BattlePassGetInfoResponse struct {
	ResponsePacket
	Protocol       Protocol
	BattlePassInfo BattlePassInfoDB
}
