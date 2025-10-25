package protos

type ClanApplicantResponse struct {
	ResponsePacket
	Protocol      Protocol
	ClanMemberDBs []ClanMemberDB
}
