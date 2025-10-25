package protos

type CafeSummonCharacterResponse struct {
	ResponsePacket
	Protocol Protocol
	CafeDB   CafeDB
	CafeDBs  []CafeDB
}
