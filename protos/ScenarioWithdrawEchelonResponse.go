package protos

type ScenarioWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol           Protocol
	SaveDataDB         StoryStrategyStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
