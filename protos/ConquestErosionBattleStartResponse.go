package protos

type ConquestErosionBattleStartResponse struct {
	ResponsePacket
	Protocol            Protocol
	ParcelResultDB      ParcelResultDB
	ConquestStageSaveDB ConquestStageSaveDB
}
