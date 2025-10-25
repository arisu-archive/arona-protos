package protos

type MiniGameCCGStagePlayDB struct {
	StageId      int64
	EnemyGroupId int64
	IsClear      bool
	RewardDBs    []MiniGameCCGStageRewardDB
	RerollPoint  int32
}
