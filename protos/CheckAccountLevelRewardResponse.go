package protos

type CheckAccountLevelRewardResponse struct {
	ResponsePacket
	Protocol              Protocol
	AccountLevelRewardIds []int64
}
