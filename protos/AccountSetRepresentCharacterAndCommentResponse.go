package protos

type AccountSetRepresentCharacterAndCommentResponse struct {
	ResponsePacket
	Protocol             Protocol
	AccountDB            AccountDB
	RepresentCharacterDB CharacterDB
}
