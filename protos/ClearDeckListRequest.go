package protos

type ClearDeckListRequest struct {
	RequestPacket
	Protocol     Protocol
	ClearDeckKey ClearDeckKey
}
