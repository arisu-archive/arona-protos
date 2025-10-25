package protos

type GuideMissionSeasonListResponse struct {
	ResponsePacket
	Protocol              Protocol
	GuideMissionSeasonDBs []GuideMissionSeasonDB
}
