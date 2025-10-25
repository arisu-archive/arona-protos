package protos

type ArenaLoginResponse struct {
	ResponsePacket
	Protocol          Protocol
	ArenaPlayerInfoDB ArenaPlayerInfoDB
}
