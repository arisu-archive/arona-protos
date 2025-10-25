package protos

type ItemLockResponse struct {
	ResponsePacket
	Protocol Protocol
	ItemDB   ItemDB
}
