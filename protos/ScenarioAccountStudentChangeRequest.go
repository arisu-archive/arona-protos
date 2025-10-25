package protos

type ScenarioAccountStudentChangeRequest struct {
	RequestPacket
	Protocol             Protocol
	AccountStudent       int64
	AccountStudentBefore int64
}
