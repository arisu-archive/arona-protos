package protos

type ConquestStepSummary struct {
	ConqueredTileCount    int64
	AllTileCount          int64
	ErosionRemainingCount int64
	HasPhaseComplete      bool
	IsErosionPhaseStart   bool
	IsStepOpen            bool
}
