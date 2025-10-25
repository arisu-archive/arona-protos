package protos

type CampaignRetreatResponse struct {
	ResponsePacket
	Protocol               Protocol
	ReleasedEchelonNumbers []int64
	ParcelResultDB         ParcelResultDB
}
