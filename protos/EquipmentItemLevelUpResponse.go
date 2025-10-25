package protos

type EquipmentItemLevelUpResponse struct {
	ResponsePacket
	Protocol          Protocol
	EquipmentDB       EquipmentDB
	AccountCurrencyDB AccountCurrencyDB
	ConsumeResultDB   ConsumeResultDB
}
