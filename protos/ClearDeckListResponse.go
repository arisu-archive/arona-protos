package protos

type ClearDeckListResponse struct {
	ResponsePacket
	Protocol     Protocol
	ClearDeckDBs []ClearDeckDB
}
