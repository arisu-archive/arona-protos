package protos

type ReceiveAccountLevelRewardResponse struct {
	ResponsePacket
	Protocol                      Protocol
	ReceivedAccountLevelRewardIds []int64
	ParcelResultDB                ParcelResultDB
}
