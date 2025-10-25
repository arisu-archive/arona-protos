package protos

type ResetableContentGetResponse struct {
	ResponsePacket
	Protocol                 Protocol
	ResetableContentValueDBs []ResetableContentValueDB
}
