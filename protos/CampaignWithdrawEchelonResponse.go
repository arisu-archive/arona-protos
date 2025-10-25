package protos

type CampaignWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol           Protocol
	SaveDataDB         CampaignMainStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
