package protos

type ConquestDisplayInfo struct {
	TriggerType  ConquestTriggerType
	Type         ConquestDisplayType
	EntityId     int64
	TileUniqueId int64
	Parameter    string
	DisplayOrder int32
	DisplayOnce  bool
}
