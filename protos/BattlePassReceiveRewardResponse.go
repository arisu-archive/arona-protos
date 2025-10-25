package protos

type BattlePassReceiveRewardResponse struct {
	ResponsePacket
	Protocol       Protocol
	BattlePassInfo BattlePassInfoDB
	ParcelResult   ParcelResultDB
}
