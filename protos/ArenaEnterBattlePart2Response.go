package protos

type ArenaEnterBattlePart2Response struct {
	ResponsePacket
	Protocol          Protocol
	ArenaBattleDB     ArenaBattleDB
	ArenaPlayerInfoDB ArenaPlayerInfoDB
	AccountCurrencyDB AccountCurrencyDB
	VictoryRewards    ParcelResultDB
	SeasonRewards     ParcelResultDB
	AllTimeRewards    ParcelResultDB
}
