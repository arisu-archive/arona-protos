package protos

type EventContentWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol           Protocol
	SaveDataDB         EventContentMainStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
