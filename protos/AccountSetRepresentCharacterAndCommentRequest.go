package protos

type AccountSetRepresentCharacterAndCommentRequest struct {
	RequestPacket
	Protocol                   Protocol
	RepresentCharacterServerId int64
	Comment                    string
}
