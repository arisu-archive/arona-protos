package protos

type OpenConditionSetResponse struct {
	ResponsePacket
	Protocol     Protocol
	ConditionDBs []OpenConditionDB
}
