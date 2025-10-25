package protos

type GearTierUpRequestDB struct {
	TargetServerId int64
	AfterTier      int64
	ReplaceInfos   []SelectTicketReplaceInfo
}
