package protos

type MiniGameDreamMakerAttendScheduleResponse struct {
	ResponsePacket
	Protocol                  Protocol
	InfoDB                    MiniGameDreamMakerInfoDB
	ParameterDBs              []MiniGameDreamMakerParameterDB
	ParcelResultDB            ParcelResultDB
	ScheduleResultId          int64
	EventContentCollectionDBs []EventContentCollectionDB
}
