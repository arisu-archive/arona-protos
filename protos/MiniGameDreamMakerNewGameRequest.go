package protos

type MiniGameDreamMakerNewGameRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	Multiplier     int64
}
