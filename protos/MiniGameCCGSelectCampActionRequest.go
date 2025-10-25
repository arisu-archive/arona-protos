package protos

type MiniGameCCGSelectCampActionRequest struct {
	RequestPacket
	Protocol        Protocol
	EventContentId  int64
	SelectedOption  MiniGameCCGCampOption
	RemoveCardDBIds []int32
}
