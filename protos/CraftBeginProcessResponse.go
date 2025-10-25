package protos

type CraftBeginProcessResponse struct {
	ResponsePacket
	Protocol    Protocol
	CraftInfoDB CraftInfoDB
}
