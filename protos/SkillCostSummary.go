package protos

type SkillCostSummary struct {
	InitialCost           float32
	CostPerFrameSnapshots CostRegenSnapshotCollection
	CostAddSnapshots      []SkillCostAddSnapshot
	CostUseSnapshots      []SkillCostUseSnapshot
}
