package protos

type UseCouponResponse struct {
	ResponsePacket
	Protocol                     Protocol
	CouponCompleteRewardReceived bool
}
