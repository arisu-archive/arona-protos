package protos

type EchelonListResponse struct {
	ResponsePacket
	Protocol       Protocol
	EchelonDBs     []EchelonDB
	ArenaEchelonDB EchelonDB
}
