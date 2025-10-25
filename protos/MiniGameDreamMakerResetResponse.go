package protos

type MiniGameDreamMakerResetResponse struct {
	ResponsePacket
	Protocol     Protocol
	InfoDB       MiniGameDreamMakerInfoDB
	ParameterDBs []MiniGameDreamMakerParameterDB
}
