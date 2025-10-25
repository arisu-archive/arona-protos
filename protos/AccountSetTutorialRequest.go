package protos

type AccountSetTutorialRequest struct {
	RequestPacket
	Protocol    Protocol
	TutorialIds []int64
}
