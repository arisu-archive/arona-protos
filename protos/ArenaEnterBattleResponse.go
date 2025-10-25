package protos

type ArenaEnterBattleResponse struct {
	ResponsePacket
	Protocol          Protocol
	AccountCurrencyDB AccountCurrencyDB
	ArenaBattleDB     ArenaBattleDB
	ArenaPlayerInfoDB ArenaPlayerInfoDB
	VictoryRewards    ParcelResultDB
	SeasonRewards     ParcelResultDB
	AllTimeRewards    ParcelResultDB
}
