package protos

type RaidBossResult struct {
	RaidDamage         RaidDamage
	EndHpRateRawValue  int64
	GroggyRateRawValue int64
	GroggyCount        int32
	SubPartsHPs        []int64
	AIPhase            int64
}
