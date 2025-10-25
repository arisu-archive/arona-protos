package protos

type ArenaSettingChangeRequest struct {
	RequestPacket
	Protocol Protocol
	MapId    int64
}
