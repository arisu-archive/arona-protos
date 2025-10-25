package protos

type CafeGetInfoRequest struct {
	RequestPacket
	Protocol        Protocol
	AccountServerId int64
}
