package protos

type EventContentBoxGachaElement struct {
	EventContentId int64
	VariationId    int64
	Round          int64
	GroupId        int64
	UniqueId       int64
	IsPrize        bool
	Rewards        []ParcelInfo
}
