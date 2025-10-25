package protos

type ClanMemberResponse struct {
	ResponsePacket
	Protocol              Protocol
	ClanDB                ClanDB
	ClanMemberDB          ClanMemberDB
	DetailedAccountInfoDB DetailedAccountInfoDB
}
