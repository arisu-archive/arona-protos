package protos

type CraftUpdateNodeLevelResponse struct {
	ResponsePacket
	Protocol          Protocol
	CraftInfoDB       CraftInfoDB
	CraftNodeDB       CraftNodeDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
