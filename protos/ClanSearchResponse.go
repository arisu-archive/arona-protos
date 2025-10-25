package protos

type ClanSearchResponse struct {
	ResponsePacket
	Protocol Protocol
	ClanDBs  []ClanDB
}
