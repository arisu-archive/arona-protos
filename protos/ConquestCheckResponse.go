package protos

type ConquestCheckResponse struct {
	ResponsePacket
	Protocol                     Protocol
	CanReceiveCalculateReward    bool
	AlarmPhaseToShow             *int32
	ParcelConsumeCumulatedAmount int64
	ConquestSummary              ConquestSummary
}
