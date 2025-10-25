package protos

type OpenConditionSetRequest struct {
	RequestPacket
	Protocol    Protocol
	ConditionDB OpenConditionDB
}
