package protos

type ConquestConquerDeployEchelonResponse struct {
	ResponsePacket
	Protocol           Protocol
	ConquestEchelonDBs []ConquestEchelonDB
	ConquestInfoDB     ConquestInfoDB
}
