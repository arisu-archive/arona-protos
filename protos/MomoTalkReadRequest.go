package protos

type MomoTalkReadRequest struct {
	RequestPacket
	Protocol               Protocol
	CharacterDBId          int64
	LastReadMessageGroupId int64
	ChosenMessageId        *int64
}
