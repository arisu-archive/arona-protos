package protos

type ConquestInfoDB struct {
	EventContentId                         int64
	EventGauge                             int32
	EventSpawnCount                        int32
	EchelonChangeCount                     int32
	TodayConquestRentCount                 int32
	TodayOperationRentCount                int32
	CumulatedConditionValue                int64
	ReceivedCalculateRewardConditionAmount int64
	AlertMassErosionId                     *int64
}
