package protos

type MiniGameDreamMakerGetInfoResponse struct {
	ResponsePacket
	Protocol                     Protocol
	InfoDB                       MiniGameDreamMakerInfoDB
	ParameterDBs                 []MiniGameDreamMakerParameterDB
	EndingDBs                    []MiniGameDreamMakerEndingDB
	EventContentCollectionDBs    []EventContentCollectionDB
	EventPointAmount             int64
	AlreadyReceivePointRewardIds []int64
}
