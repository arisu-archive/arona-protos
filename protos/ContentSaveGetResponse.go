package protos

type ContentSaveGetResponse struct {
	ResponsePacket
	Protocol             Protocol
	HasValidData         bool
	ContentSaveDB        ContentSaveDB
	EventContentChangeDB EventContentChangeDB
}
