package protos

type MiniGameDreamMakerInfoDB struct {
	EventContentId       int64
	Round                int64
	Multiplier           int64
	DayOfNumber          int64
	ActionCount          int64
	CurrentRoundEndingId int64
	EndingRewardReceived bool
	CanStartNewGame      bool
}
