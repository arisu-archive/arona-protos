package protos

type EventContentLocationAttendScheduleResponse struct {
	ResponsePacket
	Protocol                  Protocol
	EventContentLocationDB    EventContentLocationDB
	EventContentCollectionDBs []EventContentCollectionDB
	ParcelResultDB            ParcelResultDB
	ExtraRewards              []ParcelInfo
}
