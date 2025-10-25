package protos

type RaidSearchResponse struct {
	ResponsePacket
	Protocol Protocol
	RaidDBs  []RaidDB
}
