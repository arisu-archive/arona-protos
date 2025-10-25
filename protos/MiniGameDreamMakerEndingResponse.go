package protos

type MiniGameDreamMakerEndingResponse struct {
	ResponsePacket
	Protocol       Protocol
	InfoDB         MiniGameDreamMakerInfoDB
	ParameterDBs   []MiniGameDreamMakerParameterDB
	EndingDB       MiniGameDreamMakerEndingDB
	ParcelResultDB ParcelResultDB
}
