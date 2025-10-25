package protos

type MomoTalkMessageListRequest struct {
	RequestPacket
	Protocol      Protocol
	CharacterDBId int64
}
