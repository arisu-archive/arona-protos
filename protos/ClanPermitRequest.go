package protos

type ClanPermitRequest struct {
	RequestPacket
	Protocol           Protocol
	ApplicantAccountId int64
	IsPerMit           bool
}
