package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type ConquestConquerWithBattleResultResponse struct {
	ResponsePacket
	Protocol                     Protocol
	ParcelResultDB               ParcelResultDB
	ConquestTileDB               ConquestTileDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
	StepAfterBattle              int32
	DisplayParcelByRewardTag     map[flatdata.RewardTag][]ParcelInfo
}
