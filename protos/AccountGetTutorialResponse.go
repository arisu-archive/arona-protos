package protos

type AccountGetTutorialResponse struct {
	ResponsePacket
	Protocol    Protocol
	TutorialIds []int64
}
