package protos

type EventContentRetreatResponse struct {
	ResponsePacket
	Protocol               Protocol
	ReleasedEchelonNumbers []int64
	ParcelResultDB         ParcelResultDB
}
