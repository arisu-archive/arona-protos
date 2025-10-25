package protos

type ScenarioRetreatResponse struct {
	ResponsePacket
	Protocol               Protocol
	ReleasedEchelonNumbers []int64
	ParcelResultDB         ParcelResultDB
}
