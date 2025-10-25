package protos

type EventContentLocationDB struct {
	LocationId            int64
	Rank                  int64
	Exp                   int64
	ScheduleCount         int64
	ZoneVisitCharacterDBs map[int64][]VisitingCharacterDB
}
