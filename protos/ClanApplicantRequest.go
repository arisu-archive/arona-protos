package protos

type ClanApplicantRequest struct {
	RequestPacket
	Protocol Protocol
	OffSet   int64
}
