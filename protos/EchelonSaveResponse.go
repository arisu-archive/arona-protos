package protos

type EchelonSaveResponse struct {
	ResponsePacket
	Protocol  Protocol
	EchelonDB EchelonDB
}
