package protos

type BattlePassMissionSingleRewardRequest struct {
	RequestPacket
	Protocol        Protocol
	BattlePassId    int64
	MissionUniqueId int64
}
