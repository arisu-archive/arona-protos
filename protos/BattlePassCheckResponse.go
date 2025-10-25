package protos

type BattlePassCheckResponse struct {
	ResponsePacket
	Protocol            Protocol
	HasNotReceiveReward bool
	HasCompleteMission  bool
}
