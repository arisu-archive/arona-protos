package protos

type EquipmentBatchGrowthRequestDB struct {
	TargetServerId    int64
	ConsumeRequestDBs []ConsumeRequestDB
	AfterTier         int64
	AfterLevel        int64
	AfterExp          int64
	ReplaceInfos      []SelectTicketReplaceInfo
}
