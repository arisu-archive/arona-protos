package protos

type MiniGameCCGSelectRewardCardRequest struct {
	RequestPacket
	Protocol       Protocol
	EventContentId int64
	SelectedIndex  int32
	RewardIndex    int32
}
