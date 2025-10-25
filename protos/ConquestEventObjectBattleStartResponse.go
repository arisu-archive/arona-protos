package protos

type ConquestEventObjectBattleStartResponse struct {
	ResponsePacket
	Protocol            Protocol
	ParcelResultDB      ParcelResultDB
	ConquestStageSaveDB ConquestStageSaveDB
}
