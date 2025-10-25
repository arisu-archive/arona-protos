package protos

type MailReceiveResponse struct {
	ResponsePacket
	Protocol          Protocol
	MailServerIds     []int64
	ParcelResultDB    ParcelResultDB
	BattlePassInfoDBs []BattlePassInfoDB
}
