package protos

type EventContentMapMoveResponse struct {
	ResponsePacket
	Protocol                  Protocol
	SaveDataDB                EventContentMainStageSaveDB
	EchelonEntityId           int64
	StrategyObject            Strategy
	StrategyObjectParcelInfos []ParcelInfo
}
