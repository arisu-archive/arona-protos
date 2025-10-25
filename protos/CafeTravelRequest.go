package protos

type CafeTravelRequest struct {
	RequestPacket
	Protocol                 Protocol
	TargetAccountId          *int64
	CurrentVisitingAccountId *int64
}
