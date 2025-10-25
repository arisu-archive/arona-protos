package protos

type HexaDisplayInfo struct {
	Type            HexaDisplayType
	EntityId        int64
	UniqueId        int64
	Location        HexLocation
	Parameter       int64
	StageRewardInfo StrategyClearRewardInfo
}
