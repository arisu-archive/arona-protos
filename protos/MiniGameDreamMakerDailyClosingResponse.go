package protos

type MiniGameDreamMakerDailyClosingResponse struct {
	ResponsePacket
	Protocol                     Protocol
	InfoDB                       MiniGameDreamMakerInfoDB
	ParameterDBs                 []MiniGameDreamMakerParameterDB
	ParcelResultDB               ParcelResultDB
	EventPointAmount             int64
	AlreadyReceivePointRewardIds []int64
}
