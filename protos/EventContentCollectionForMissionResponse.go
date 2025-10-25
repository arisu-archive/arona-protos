package protos

type EventContentCollectionForMissionResponse struct {
	ResponsePacket
	Protocol                  Protocol
	EventContentCollectionDBs []EventContentCollectionDB
}
