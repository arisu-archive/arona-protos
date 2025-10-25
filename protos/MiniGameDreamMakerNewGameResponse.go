package protos

type MiniGameDreamMakerNewGameResponse struct {
	ResponsePacket
	Protocol     Protocol
	InfoDB       MiniGameDreamMakerInfoDB
	ParameterDBs []MiniGameDreamMakerParameterDB
}
