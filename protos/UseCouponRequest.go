package protos

type UseCouponRequest struct {
	RequestPacket
	Protocol     Protocol
	CouponSerial string
}
