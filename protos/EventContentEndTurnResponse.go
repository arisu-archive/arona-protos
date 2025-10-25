package protos

type EventContentEndTurnResponse struct {
	ResponsePacket
	Protocol          Protocol
	SaveDataDB        EventContentMainStageSaveDB
	AccountCurrencyDB AccountCurrencyDB
}
