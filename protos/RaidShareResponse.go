package protos

type RaidShareResponse struct {
	ResponsePacket
	Protocol Protocol
	RaidDB   RaidDB
}
