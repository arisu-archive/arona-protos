package protos

type RaidDetailResponse struct {
	ResponsePacket
	Protocol                      Protocol
	RaidDetailDB                  RaidDetailDB
	ParticipateCharacterServerIds []int64
}
