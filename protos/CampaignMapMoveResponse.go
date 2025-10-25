package protos

type CampaignMapMoveResponse struct {
	ResponsePacket
	Protocol                  Protocol
	SaveDataDB                CampaignMainStageSaveDB
	EchelonEntityId           int64
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
